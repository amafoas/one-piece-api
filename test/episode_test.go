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

var testEpisode = models.Episode{
	ID:              "20",
	Title:           "test title",
	Release:         "October 20, 1999",
	RemasterRelease: "",
	Characters:      []string{"first", "second"},
	Season:          3,
	Episode:         23,
	Locations:       []string{"first location", "second location"},
	Opening:         "Opening song",
}

type EpisodeTestSuite struct {
	suite.Suite
}

func (suite *EpisodeTestSuite) Test1InsertEpisode() {
	episodeJSON, err := json.Marshal(testEpisode)
	require.NoError(suite.T(), err)

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/episode", bytes.NewBuffer(episodeJSON))
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
}

func (suite *EpisodeTestSuite) Test2GetEpisodeByID() {
	route := fmt.Sprintf("/episode/%s", testEpisode.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseEpisode models.Episode
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseEpisode)
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), testEpisode, reponseEpisode)
}

func (suite *EpisodeTestSuite) Test3UpdateEpisode() {
	route := fmt.Sprintf("/episode/%s", testEpisode.ID)

	payload := []byte(`{
		"title": "Updated title",
		"season": 25
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

	// Send a new GET request to retrieve the updated episode
	req, err = http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Reset the ResponseRecorder for the new request
	recorder = httptest.NewRecorder()
	// Send the test HTTP request to the router again
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseEpisode models.Episode
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseEpisode)
	require.NoError(suite.T(), err)

	updatedEpisode := testEpisode
	updatedEpisode.Title = "Updated title"
	updatedEpisode.Season = 25

	assert.Equal(suite.T(), updatedEpisode, reponseEpisode)
}

func (suite *EpisodeTestSuite) Test4DeleteEpisode() {
	route := fmt.Sprintf("/episode/%s", testEpisode.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("DELETE", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	// Try to get the deleted episode
	getReq, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, getReq)

	// Check that the episode has been removed
	assert.Equal(suite.T(), http.StatusNotFound, recorder.Code)
}

func TestEpisodeSuite(t *testing.T) {
	suite.Run(t, new(EpisodeTestSuite))
}
