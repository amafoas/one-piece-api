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

var testCrew = models.Crew{
	ID:              "test-pirate-crew",
	Name:            "Test pirate crew",
	RomanizedName:   "Testo piratsu",
	FirstAppearance: []string{"Chapter 10", "Episode 30"},
	Captain:         "Guybrush Threepwood",
	TotalBounty:     "9,999,999,999",
	MainShip:        "Super Test Ship",
	Members:         []string{"Guybrush Threepwood", "Elaine Marley", "LeChuck"},
	Allies:          []string{"Super pirate allies crew"},
}

func configCrewRoutesTest(db *mongo.Database) {
	log.Println("Crew routes test setup")
	routers.ConfigureCrewRoutes(router)

	collection := db.Collection("crews")

	_, err := collection.InsertOne(context.TODO(), testCrew)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Test crew inserted in the database.")
}

func TestGetCrewByID(t *testing.T) {
	path := fmt.Sprintf("/crew/%s", testCrew.ID)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Incorrect status code")

	expectedBody, err := json.Marshal(testCrew)
	if err != nil {
		t.Error("Error while trying to marshal")
	}
	assert.Equal(t, expectedBody, response.Body.Bytes(), "Wrong Answer Body")
}

func TestGetCrewNotFound(t *testing.T) {
	path := fmt.Sprintf("/crew/%s", "invalid")
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNotFound, response.Code, "Incorrect status code")
}
