package handlers

import (
	"net/http"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CharacterByID(c *gin.Context) {
	id := c.Param("id")

	character, err := repositories.GetCharacterByField("_id", id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, character)
}

func CharacterByFruit(c *gin.Context) {
	fruit := c.Param("fruit")

	character, err := repositories.GetCharacterByField("devil_fruit_id", fruit)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, character)
}
