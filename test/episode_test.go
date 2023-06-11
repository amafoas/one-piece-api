package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"one-piece-api/api/models"
	"one-piece-api/api/routers"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var testEpisode = models.Episode{
	Title:           "test title",
	Release:         "October 20, 1999",
	RemasterRelease: "",
	Characters:      []string{"first", "second"},
	Season:          3,
	Episode:         23,
	Locations:       []string{"first location", "second location"},
	Opening:         "Opening song",
}

func configEpisodeRoutesTest(db *mongo.Database) {
	log.Println("Episodes routes test setup")
	routers.ConfigureEpisodeRoutes(router)

	collection := db.Collection("episodes")

	_, err := collection.InsertOne(context.TODO(), testEpisode)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Test episode inserted in the database.")
}

func TestGetEpisodeByNumber(t *testing.T) {
	path := fmt.Sprintf("/episode/%v", testEpisode.Episode)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Incorrect status code")

	expectedBody, err := json.Marshal(testEpisode)
	if err != nil {
		t.Error("Error while trying to marshal")
	}
	assert.Equal(t, expectedBody, response.Body.Bytes(), "Wrong Answer Body")
}

func TestGetEpisodeNotFound(t *testing.T) {
	path := fmt.Sprintf("/episode/%v", 0)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNotFound, response.Code, "Incorrect status code")
}
