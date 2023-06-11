package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"one-piece-api/api/models"
	"one-piece-api/api/routers"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var testDevilFruit = models.DevilFruit{
	ID:              "test-test-no-mi",
	Name:            "Test Test No Mi",
	Type:            "Logia",
	Meaning:         "Sound of test passing",
	FirstApparition: []string{"Chapter 999", "Episode 444"},
	FirstUsage:      []string{"Chapter 444", "Episode 999"},
	CurrentUser:     "Keanu Reeves",
	PreviousUser:    "",
}

func configDevilFruitRoutesTest(db *mongo.Database) {
	log.Println("Devil fruits routes test setup")
	routers.ConfigureDevilFruitsRoutes(router)

	collection := db.Collection("devil_fruits")

	_, err := collection.InsertOne(context.TODO(), testDevilFruit)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Test fruit inserted in the database.")
}

func TestGetDevilFruitByID(t *testing.T) {
	path := fmt.Sprintf("/devil-fruit/%s", testDevilFruit.ID)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Incorrect status code")

	expectedBody, err := json.Marshal(testDevilFruit)
	if err != nil {
		t.Error("Error while trying to marshal")
	}
	assert.Equal(t, expectedBody, response.Body.Bytes(), "Wrong Answer Body")
}

func TestGetDevilFruitNotFound(t *testing.T) {
	path := fmt.Sprintf("/devil-fruit/%s", "none-none-no-mi")
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNotFound, response.Code, "Incorrect status code")
}
