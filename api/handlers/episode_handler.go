package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"one-piece-api/api/models"
	"one-piece-api/api/repositories"
)

func CreateEpisode(c *gin.Context) {
	var newEpisode models.Episode
	if err := c.ShouldBindJSON(&newEpisode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	episodeRepository := repositories.NewEpisodeRepository(db)
	err = episodeRepository.Create(newEpisode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}

func FindEpisodeByID(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	episodeRepository := repositories.NewEpisodeRepository(db)

	var episode models.Episode
	err = episodeRepository.FindByID(id, &episode)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, episode)
}

func UpdateEpisode(c *gin.Context) {
	id := c.Param("id")

	var updatedEpisode map[string]interface{}
	if err := c.ShouldBindJSON(&updatedEpisode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	episodeRepository := repositories.NewEpisodeRepository(db)
	err = episodeRepository.Update(id, updatedEpisode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func DeleteEpisode(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	episodeRepository := repositories.NewEpisodeRepository(db)
	err = episodeRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
