package handlers

import (
	"net/http"
	"one-piece-api/api/models"
	"one-piece-api/api/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCrew(c *gin.Context) {
	var newCrew models.Crew
	if err := c.ShouldBindJSON(&newCrew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	crewRepository := repositories.NewCrewRepository(db)
	err = crewRepository.Create(newCrew)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Document created successfully"})
}

func FindCrewByID(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	crewRepository := repositories.NewCrewRepository(db)

	var crew models.Crew
	err = crewRepository.FindByID(id, &crew)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, crew)
}

func UpdateCrew(c *gin.Context) {
	id := c.Param("id")

	var updatedCrew map[string]interface{}
	if err := c.ShouldBindJSON(&updatedCrew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	crewRepository := repositories.NewCrewRepository(db)
	err = crewRepository.Update(id, updatedCrew)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func DeleteCrew(c *gin.Context) {
	id := c.Param("id")

	db, err := repositories.GetDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	crewRepository := repositories.NewCrewRepository(db)
	err = crewRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
