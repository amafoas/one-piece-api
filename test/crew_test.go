package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"one-piece-api/api/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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

type CrewTestSuite struct {
	suite.Suite
}

func (suite *CrewTestSuite) Test1InsertCrew() {
	crewJSON, err := json.Marshal(testCrew)
	require.NoError(suite.T(), err)

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/crew", bytes.NewBuffer(crewJSON))
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
}

func (suite *CrewTestSuite) Test2GetCrewByID() {
	route := fmt.Sprintf("/crew/%s", testCrew.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseCrew models.Crew
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseCrew)
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), testCrew, reponseCrew)
}

func (suite *CrewTestSuite) Test3UpdateCrew() {
	route := fmt.Sprintf("/crew/%s", testCrew.ID)

	payload := []byte(`{
		"name": "Updated name",
		"captain": "Updated captain"
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

	// Send a new GET request to retrieve the updated crew
	req, err = http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Reset the ResponseRecorder for the new request
	recorder = httptest.NewRecorder()
	// Send the test HTTP request to the router again
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseCrew models.Crew
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseCrew)
	require.NoError(suite.T(), err)

	updatedCrew := testCrew
	updatedCrew.Name = "Updated name"
	updatedCrew.Captain = "Updated captain"

	assert.Equal(suite.T(), updatedCrew, reponseCrew)
}

func (suite *CrewTestSuite) Test4DeleteCrew() {
	route := fmt.Sprintf("/crew/%v", testCrew.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("DELETE", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	// Try to get the deleted crew
	getReq, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, getReq)

	// Check that the crew has been removed
	assert.Equal(suite.T(), http.StatusNotFound, recorder.Code)
}

func TestCrewSuite(t *testing.T) {
	suite.Run(t, new(CrewTestSuite))
}
