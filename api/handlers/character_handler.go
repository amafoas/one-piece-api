package handlers

import (
	"log"
	"net/http"
	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCharacter(c *gin.Context) {
	var newCharacter models.Character
	if err := c.ShouldBindJSON(&newCharacter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(newCharacter)

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	characterRepository := repositories.NewCharacterRepository(db)
	err = characterRepository.Create(newCharacter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}

func FindCharacterByID(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	characterRepository := repositories.NewCharacterRepository(db)

	var character models.Character
	err = characterRepository.FindByID(id, &character)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, character)
}

func UpdateCharacter(c *gin.Context) {
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

	characterRepository := repositories.NewCharacterRepository(db)
	err = characterRepository.Update(id, updatedChapter)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func DeleteCharacter(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	characterRepository := repositories.NewCharacterRepository(db)
	err = characterRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
