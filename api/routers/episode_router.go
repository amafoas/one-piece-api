package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureEpisodeRoutes(r *gin.Engine) {
	r.GET("/episode/:num", handlers.EpisodeByNumber)
}
