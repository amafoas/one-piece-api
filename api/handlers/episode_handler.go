package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"one-piece-api/api/repositories"
)

func EpisodeByNumber(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		log.Println(err)
	}

	episode, err := repositories.GetEpisodeByNumber(num)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, episode)
}
