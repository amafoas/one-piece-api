package routers

import (
	"github.com/gin-gonic/gin"

	"one-piece-api/api/handlers"
)

func ConfigureCharacterRoutes(r *gin.Engine) {
	character := r.Group("/character")
	{
		character.GET("/:id", handlers.CharacterByID)
		character.GET("/fruit/:fruit", handlers.CharacterByFruit)
	}
}
