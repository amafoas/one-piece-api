package handlers

import (
	"net/http"
	"reflect"

	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
)

type ChapterHandler struct {
	*BaseHandler
}

func NewChapterHandler(c *gin.Context) *ChapterHandler {
	repo, err := repositories.NewChapterRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return nil
	}
	model := reflect.TypeOf(models.Chapter{})

	return &ChapterHandler{
		BaseHandler: NewHandler(repo, model),
	}
}

func CreateChapter(c *gin.Context) {
	handler := NewChapterHandler(c)
	handler.CreateEntity(c)
}

func FindChapterByID(c *gin.Context) {
	handler := NewChapterHandler(c)
	handler.FindEntityByID(c)
}

func UpdateChapter(c *gin.Context) {
	handler := NewChapterHandler(c)
	handler.UpdateEntity(c)
}

func DeleteChapter(c *gin.Context) {
	handler := NewChapterHandler(c)
	handler.DeleteEntity(c)
}
