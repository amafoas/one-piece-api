package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureCrewRoutes(r *gin.Engine) {
	r.GET("/crew/:id", handlers.CrewByID)
}
