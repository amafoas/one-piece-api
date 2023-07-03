package handlers

import (
	"net/http"
	"reflect"

	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
)

type DevilFruitHandler struct {
	*BaseHandler
}

func NewDevilFruitHandler(c *gin.Context) *DevilFruitHandler {
	repo, err := repositories.NewDevilFruitRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return nil
	}
	model := reflect.TypeOf(models.DevilFruit{})

	return &DevilFruitHandler{
		BaseHandler: NewHandler(repo, model),
	}
}

func CreateDevilFruit(c *gin.Context) {
	handler := NewDevilFruitHandler(c)
	handler.CreateEntity(c)
}

func FindDevilFruitByID(c *gin.Context) {
	handler := NewDevilFruitHandler(c)
	handler.FindEntityByID(c)
}

func UpdateDevilFruit(c *gin.Context) {
	handler := NewDevilFruitHandler(c)
	handler.UpdateEntity(c)
}

func DeleteDevilFruit(c *gin.Context) {
	handler := NewDevilFruitHandler(c)
	handler.DeleteEntity(c)
}
