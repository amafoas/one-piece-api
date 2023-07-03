package handlers

import (
	"net/http"
	"reflect"

	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
)

type CrewHandler struct {
	*BaseHandler
}

func NewCrewHandler(c *gin.Context) *CrewHandler {
	repo, err := repositories.NewCrewRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return nil
	}
	model := reflect.TypeOf(models.Crew{})

	return &CrewHandler{
		BaseHandler: NewHandler(repo, model),
	}
}

func CreateCrew(c *gin.Context) {
	handler := NewCrewHandler(c)
	handler.CreateEntity(c)
}

func FindCrewByID(c *gin.Context) {
	handler := NewCrewHandler(c)
	handler.FindEntityByID(c)
}

func UpdateCrew(c *gin.Context) {
	handler := NewCrewHandler(c)
	handler.UpdateEntity(c)
}

func DeleteCrew(c *gin.Context) {
	handler := NewCrewHandler(c)
	handler.DeleteEntity(c)
}
