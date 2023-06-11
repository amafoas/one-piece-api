package handlers

import (
	"net/http"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CrewByID(c *gin.Context) {
	id := c.Param("id")

	crew, err := repositories.GetCrewByID(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Server error"})
		return
	}

	// crewJSON, err := json.Marshal(crew)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(crew)
	c.JSON(http.StatusOK, crew)
}
