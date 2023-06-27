package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureChapterRoutes(r *gin.Engine) {
	r.GET("/chapter/:id", handlers.FindChapterByID)

	r.POST("/chapter", handlers.CreateChapter)

	r.PUT("/chapter/:id", handlers.UpdateChapter)

	r.DELETE("/chapter/:id", handlers.DeleteChapter)
}
