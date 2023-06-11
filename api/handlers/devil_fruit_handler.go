package handlers

import (
	"net/http"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DevilFruitByID(c *gin.Context) {
	id := c.Param("id")

	fruit, err := repositories.GetDevilFruitByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, fruit)
}
