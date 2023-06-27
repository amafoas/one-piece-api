package handlers

import (
	"net/http"
	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateDevilFruit(c *gin.Context) {
	var newFruit models.DevilFruit
	if err := c.ShouldBindJSON(&newFruit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	devilFruitRepository := repositories.NewDevilFruitRepository(db)
	err = devilFruitRepository.Create(newFruit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}

func FindDevilFruitByID(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	devilFruitRepository := repositories.NewDevilFruitRepository(db)

	var devilFruit models.DevilFruit
	err = devilFruitRepository.FindByID(id, &devilFruit)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, devilFruit)
}

func UpdateDevilFruit(c *gin.Context) {
	id := c.Param("id")

	var updatedDevilFruit map[string]interface{}
	if err := c.ShouldBindJSON(&updatedDevilFruit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	devilFruitRepository := repositories.NewDevilFruitRepository(db)
	err = devilFruitRepository.Update(id, updatedDevilFruit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func DeleteDevilFruit(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	devilFruitRepository := repositories.NewDevilFruitRepository(db)
	err = devilFruitRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
