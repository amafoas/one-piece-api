package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"one-piece-api/api/routers"
	"one-piece-api/utils"
)

func main() {
	// loading .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// server conection
	r := gin.Default()

	routers.ConfigureCharacterRoutes(r)
	routers.ConfigureDevilFruitsRoutes(r)
	routers.ConfigureCrewRoutes(r)
	routers.ConfigureChapterRoutes(r)
	routers.ConfigureEpisodeRoutes(r)

	PORT := utils.GetEnvVariable("PORT")
	r.Run(":" + PORT)
}
