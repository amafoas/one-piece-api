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

type DevilFruitTestSuite struct {
	suite.Suite
}

func (suite *DevilFruitTestSuite) Test1InsertDevilFruit() {
	devilFruitJSON, err := json.Marshal(testDevilFruit)
	require.NoError(suite.T(), err)

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/devil-fruit", bytes.NewBuffer(devilFruitJSON))
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
}

func (suite *DevilFruitTestSuite) Test2GetDevilFruitByID() {
	route := fmt.Sprintf("/devil-fruit/%v", testDevilFruit.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseDevilFruit models.DevilFruit
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseDevilFruit)
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), testDevilFruit, reponseDevilFruit)
}

func (suite *DevilFruitTestSuite) Test3UpdateDevilFruit() {
	route := fmt.Sprintf("/devil-fruit/%v", testDevilFruit.ID)

	payload := []byte(`{
		"name": "Updated name",
		"type": "Updated type"
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

	// Send a new GET request to retrieve the updated devil fruit
	req, err = http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Reset the ResponseRecorder for the new request
	recorder = httptest.NewRecorder()
	// Send the test HTTP request to the router again
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseDevilFruit models.DevilFruit
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseDevilFruit)
	require.NoError(suite.T(), err)

	updatedDevilFruit := testDevilFruit
	updatedDevilFruit.Name = "Updated name"
	updatedDevilFruit.Type = "Updated type"

	assert.Equal(suite.T(), updatedDevilFruit, reponseDevilFruit)
}

func (suite *DevilFruitTestSuite) Test4DeleteDevilFruit() {
	route := fmt.Sprintf("/devil-fruit/%v", testDevilFruit.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("DELETE", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	// Try to get the deleted devil fruit
	getReq, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, getReq)

	// Check that the devil fruit has been removed
	assert.Equal(suite.T(), http.StatusNotFound, recorder.Code)
}

func TestDevilFruitSuite(t *testing.T) {
	suite.Run(t, new(DevilFruitTestSuite))
}
