package routers

import (
	"one-piece-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigureDevilFruitsRoutes(r *gin.Engine) {
	r.GET("/devil-fruit/:id", handlers.DevilFruitByID)
}
