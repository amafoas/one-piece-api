package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureCrewRoutes(r *gin.Engine) {
	r.GET("/crew/:id", handlers.FindCrewByID)

	r.POST("/crew", handlers.CreateCrew)

	r.PUT("/crew/:id", handlers.UpdateCrew)

	r.DELETE("/crew/:id", handlers.DeleteCrew)
}
