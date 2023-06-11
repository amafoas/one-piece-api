package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureChapterRoutes(r *gin.Engine) {
	r.GET("/chapter/:num", handlers.ChapterByNumber)
}
