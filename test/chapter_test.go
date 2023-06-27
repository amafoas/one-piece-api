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

var testChapter = models.Chapter{
	ID:      "33",
	Title:   "test chapter",
	Volume:  50,
	Chapter: 33,
	Pages:   10,
	Release: "May, 10th",
}

type ChapterTestSuite struct {
	suite.Suite
}

func (suite *ChapterTestSuite) Test1InsertChapter() {
	chapterJSON, err := json.Marshal(testChapter)
	require.NoError(suite.T(), err)

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/chapter", bytes.NewBuffer(chapterJSON))
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
}

func (suite *ChapterTestSuite) Test2GetChapterByID() {
	route := fmt.Sprintf("/chapter/%s", testChapter.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseChapter models.Chapter
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseChapter)
	require.NoError(suite.T(), err)

	assert.Equal(suite.T(), testChapter, reponseChapter)
}

func (suite *ChapterTestSuite) Test3UpdateChapter() {
	route := fmt.Sprintf("/chapter/%s", testChapter.ID)

	payload := []byte(`{
		"title": "Updated title",
		"pages": 25
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

	// Send a new GET request to retrieve the updated chapter
	req, err = http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)
	// Reset the ResponseRecorder for the new request
	recorder = httptest.NewRecorder()
	// Send the test HTTP request to the router again
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	var reponseChapter models.Chapter
	err = json.Unmarshal(recorder.Body.Bytes(), &reponseChapter)
	require.NoError(suite.T(), err)

	updatedChapter := testChapter
	updatedChapter.Title = "Updated title"
	updatedChapter.Pages = 25

	assert.Equal(suite.T(), updatedChapter, reponseChapter)
}

func (suite *ChapterTestSuite) Test4DeleteChapter() {
	route := fmt.Sprintf("/chapter/%s", testChapter.ID)

	// Create a test HTTP request
	req, err := http.NewRequest("DELETE", route, nil)
	require.NoError(suite.T(), err)
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	// Try to get the deleted chapter
	getReq, err := http.NewRequest("GET", route, nil)
	require.NoError(suite.T(), err)

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, getReq)

	// Check that the chapter has been removed
	assert.Equal(suite.T(), http.StatusNotFound, recorder.Code)
}

func TestChapterSuite(t *testing.T) {
	suite.Run(t, new(ChapterTestSuite))
}
