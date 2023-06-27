package routers

import (
	"github.com/gin-gonic/gin"

	"one-piece-api/api/handlers"
)

func ConfigureCharacterRoutes(r *gin.Engine) {
	r.GET("/character/:id", handlers.FindCharacterByID)

	r.POST("/character", handlers.CreateCharacter)

	r.PUT("/character/:id", handlers.UpdateCharacter)

	r.DELETE("/character/:id", handlers.DeleteCharacter)
}
