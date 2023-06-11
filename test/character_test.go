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

var testCharacter = models.Character{
	ID:               "benson-dunwoody",
	Name:             "Benson Dunwoody",
	Age:              45,
	Status:           "Alive",
	DevilFruit:       "Mochi Mochi No Mi",
	DevilFruitID:     "mochi-mochi-no-mi",
	Debut:            []string{},
	MainAffiliation:  "Benson's Family",
	OtherAffiliation: []string{""},
	Occupations:      "Park Owner",
	Origin:           "Apartment 1635",
	Race:             "Gumball Machine",
	Bounty:           "100",
	Birthday:         "December 18, 1989",
	Height:           "193 cm",
}

func configCharacterRoutesTest(db *mongo.Database) {
	log.Println("Character routes test setup")
	routers.ConfigureCharacterRoutes(router)

	collection := db.Collection("characters")

	_, err := collection.InsertOne(context.TODO(), testCharacter)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Test character inserted in the database.")
}
func TestGetCharacterByID(t *testing.T) {
	path := fmt.Sprintf("/character/%s", testCharacter.ID)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Incorrect status code")

	expectedBody, err := json.Marshal(testCharacter)
	if err != nil {
		t.Error("Error while trying to marshal")
	}
	assert.Equal(t, expectedBody, response.Body.Bytes(), "Wrong Answer Body")
}

func TestGetCharacterNotFound(t *testing.T) {
	path := fmt.Sprintf("/character/%s", "null")
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNotFound, response.Code, "Incorrect status code")
}

func TestGetCharacterByFruit(t *testing.T) {
	path := fmt.Sprintf("/character/fruit/%s", testCharacter.DevilFruitID)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Incorrect status code")

	expectedBody, err := json.Marshal(testCharacter)
	if err != nil {
		t.Error("Error while trying to marshal")
	}
	assert.Equal(t, expectedBody, response.Body.Bytes(), "Wrong Answer Body")
}
