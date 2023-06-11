package test

import (
	"context"
	"log"
	"one-piece-api/api/repositories"
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
	db.Drop(context.TODO())

	router = gin.Default()

	configChapterRoutesTest(db)
	configEpisodeRoutesTest(db)
	configDevilFruitRoutesTest(db)
	configCrewRoutesTest(db)
	configCharacterRoutesTest(db)

	exitCode := m.Run()
	os.Exit(exitCode)
}
