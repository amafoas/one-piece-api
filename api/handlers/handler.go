package handlers

import (
	"net/http"
	"reflect"

	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseHandler struct {
	repository repositories.Repository
	modelType  reflect.Type
}

func NewHandler(repository repositories.Repository, modelType reflect.Type) *BaseHandler {
	return &BaseHandler{
		repository: repository,
		modelType:  modelType,
	}
}

func (h *BaseHandler) CreateEntity(c *gin.Context) {
	newEntity := reflect.New(h.modelType).Interface()
	if err := c.ShouldBindJSON(newEntity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repository.Create(newEntity)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}

func (h *BaseHandler) FindEntityByID(c *gin.Context) {
	id := c.Param("id")

	newEntity := reflect.New(h.modelType).Interface()

	err := h.repository.FindByID(id, newEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, newEntity)
}

func (h *BaseHandler) UpdateEntity(c *gin.Context) {
	id := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repository.Update(id, updates)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func (h *BaseHandler) DeleteEntity(c *gin.Context) {
	id := c.Param("id")

	err := h.repository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
