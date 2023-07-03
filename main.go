package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"one-piece-api/api/routers"
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

	PORT := os.Getenv("PORT")
	r.Run(":" + PORT)
}
