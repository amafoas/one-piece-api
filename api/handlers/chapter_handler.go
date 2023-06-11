package handlers

import (
	"log"
	"net/http"
	"one-piece-api/api/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ChapterByNumber(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		log.Println(err)
	}

	chapter, err := repositories.GetChapterByNumber(num)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, chapter)
}
