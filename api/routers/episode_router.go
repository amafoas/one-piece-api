package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureEpisodeRoutes(r *gin.Engine) {
	r.GET("/episode/:id", handlers.FindEpisodeByID)

	r.POST("/episode", handlers.CreateEpisode)

	r.PUT("/episode/:id", handlers.UpdateEpisode)

	r.DELETE("/episode/:id", handlers.DeleteEpisode)
}
