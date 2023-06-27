package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"one-piece-api/api/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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

type CharacterTestSuite struct {
	suite.Suite
}

func (suite *CharacterTestSuite) Test1InsertCharacter() {
	characterJSON, err := json.Marshal(testCharacter)
	require.NoError(suite.T(), err)

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/character", bytes.NewBuffer(characterJSON))
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
}

func (suite *CharacterTestSuite) Test2GetCharacterByID() {
	route := fmt.Sprintf("/character/%s", testCharacter.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseCharacter models.Character
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseCharacter)
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), testCharacter, reponseCharacter)
}

func (suite *CharacterTestSuite) Test3UpdateCharacter() {
	route := fmt.Sprintf("/character/%s", testCharacter.ID)

	payload := []byte(`{
		"name": "Updated name",
		"origin": "Updated origin"
		}`)
	// Create a test HTTP request
	req, err := http.NewRequest("PUT", route, bytes.NewBuffer(payload))
	require.NoError(suite.T(), err)
	req.Header.Set("Content-Type", "application/json")
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	// Send a new GET request to retrieve the updated character
	req, err = http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Reset the ResponseRecorder for the new request
	recorder = httptest.NewRecorder()
	// Send the test HTTP request to the router again
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var responseCharacter models.Character
	err = json.Unmarshal(recorder.Body.Bytes(), &responseCharacter)
	require.NoError(suite.T(), err)

	updatedCharacter := testCharacter
	updatedCharacter.Name = "Updated name"
	updatedCharacter.Origin = "Updated origin"

	assert.Equal(suite.T(), updatedCharacter, responseCharacter)
}

func (suite *CharacterTestSuite) Test4DeleteCharacter() {
	route := fmt.Sprintf("/character/%s", testCharacter.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("DELETE", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	// Try to get the deleted character
	getReq, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, getReq)

	// Check that the character has been removed
	assert.Equal(suite.T(), http.StatusNotFound, recorder.Code)
}

func TestCharacterSuite(t *testing.T) {
	suite.Run(t, new(CharacterTestSuite))
}
