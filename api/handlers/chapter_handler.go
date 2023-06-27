package handlers

import (
	"net/http"
	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateChapter(c *gin.Context) {
	var newChapter models.Chapter
	if err := c.ShouldBindJSON(&newChapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	chapterRepository := repositories.NewChapterRepository(db)
	err = chapterRepository.Create(newChapter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}

func FindChapterByID(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	chapterRepository := repositories.NewChapterRepository(db)

	var chapter models.Chapter
	err = chapterRepository.FindByID(id, &chapter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, chapter)
}

func UpdateChapter(c *gin.Context) {
	id := c.Param("id")

	var updatedChapter map[string]interface{}
	if err := c.ShouldBindJSON(&updatedChapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	chapterRepository := repositories.NewChapterRepository(db)
	err = chapterRepository.Update(id, updatedChapter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func DeleteChapter(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	chapterRepository := repositories.NewChapterRepository(db)
	err = chapterRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
