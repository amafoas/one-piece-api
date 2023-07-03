package unit

import (
	"context"
	"log"
	"os"
	"testing"

	"one-piece-api/api/repositories"

	"github.com/joho/godotenv"
)

var repo *repositories.BaseRepository

func TestMain(m *testing.M) {
	err := godotenv.Load("../config/.env.test")
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

	// Base repository test
	repo = &repositories.BaseRepository{
		Collection: db.Collection("base_repository"),
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}
