package handlers

import (
	"net/http"
	"reflect"

	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
)

type EpisodeHandler struct {
	*BaseHandler
}

func NewEpisodeHandler(c *gin.Context) *EpisodeHandler {
	repo, err := repositories.NewEpisodeRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return nil
	}
	model := reflect.TypeOf(models.Episode{})

	return &EpisodeHandler{
		BaseHandler: NewHandler(repo, model),
	}
}

func CreateEpisode(c *gin.Context) {
	handler := NewEpisodeHandler(c)
	handler.CreateEntity(c)
}

func FindEpisodeByID(c *gin.Context) {
	handler := NewEpisodeHandler(c)
	handler.FindEntityByID(c)
}

func UpdateEpisode(c *gin.Context) {
	handler := NewEpisodeHandler(c)
	handler.UpdateEntity(c)
}

func DeleteEpisode(c *gin.Context) {
	handler := NewEpisodeHandler(c)
	handler.DeleteEntity(c)
}
