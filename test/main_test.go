package test

import (
	"context"
	"log"
	"one-piece-api/api/repositories"
	"one-piece-api/api/routers"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatal(err)
	}

	db, err := repositories.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Drop(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	router = gin.Default()

	routers.ConfigureChapterRoutes(router)
	routers.ConfigureCharacterRoutes(router)
	routers.ConfigureCrewRoutes(router)
	routers.ConfigureDevilFruitsRoutes(router)
	routers.ConfigureEpisodeRoutes(router)

	exitCode := m.Run()
	os.Exit(exitCode)
}
