package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureDevilFruitsRoutes(r *gin.Engine) {
	r.GET("/devil-fruit/:id", handlers.FindDevilFruitByID)

	r.POST("/devil-fruit", handlers.CreateDevilFruit)

	r.PUT("/devil-fruit/:id", handlers.UpdateDevilFruit)

	r.DELETE("/devil-fruit/:id", handlers.DeleteDevilFruit)
}
