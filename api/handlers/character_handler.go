package handlers

import (
	"net/http"
	"reflect"

	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	*BaseHandler
}

func NewCharacterHandler(c *gin.Context) *CharacterHandler {
	repo, err := repositories.NewCharacterRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return nil
	}
	model := reflect.TypeOf(models.Character{})

	return &CharacterHandler{
		BaseHandler: NewHandler(repo, model),
	}
}

func CreateCharacter(c *gin.Context) {
	handler := NewCharacterHandler(c)
	handler.CreateEntity(c)
}

func FindCharacterByID(c *gin.Context) {
	handler := NewCharacterHandler(c)
	handler.FindEntityByID(c)
}

func UpdateCharacter(c *gin.Context) {
	handler := NewCharacterHandler(c)
	handler.UpdateEntity(c)
}

func DeleteCharacter(c *gin.Context) {
	handler := NewCharacterHandler(c)
	handler.DeleteEntity(c)
}
