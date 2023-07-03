package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"one-piece-api/api/repositories"
	"one-piece-api/api/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var router *gin.Engine

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

	router = gin.Default()

	routers.ConfigureChapterRoutes(router)
	routers.ConfigureCharacterRoutes(router)
	routers.ConfigureCrewRoutes(router)
	routers.ConfigureDevilFruitsRoutes(router)
	routers.ConfigureEpisodeRoutes(router)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func makeRequest(method string, route string, requestBody interface{}) (*httptest.ResponseRecorder, error) {
	// Serialize the request body in JSON format
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	// Create an HTTP request
	req, err := http.NewRequest(method, route, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// Create a ResponseRecorder to capture the HTTP response
	recorder := httptest.NewRecorder()
	// Send the test HTTP request to the router
	router.ServeHTTP(recorder, req)

	return recorder, nil
}

type TestData struct {
	Route        string
	RequestBody  interface{}
	ExpectedBody interface{}
}

func performInsertTest(t *testing.T, config TestData) {
	recorder, err := makeRequest("POST", config.Route, config.RequestBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}

func performGetTest(t *testing.T, config TestData) {
	recorder, err := makeRequest("GET", config.Route, nil)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, recorder.Code)

	// vetify body
	responseType := reflect.TypeOf(config.ExpectedBody)
	response := reflect.New(responseType).Interface()
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, config.ExpectedBody, reflect.ValueOf(response).Elem().Interface())
}

func performUpdateTest(t *testing.T, config TestData) {
	putRecorder, err := makeRequest("PUT", config.Route, config.RequestBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, putRecorder.Code)

	// recover updated obj
	getRecorder, err := makeRequest("GET", config.Route, nil)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, getRecorder.Code)

	// vetify body
	responseType := reflect.TypeOf(config.ExpectedBody)
	response := reflect.New(responseType).Interface()
	err = json.Unmarshal(getRecorder.Body.Bytes(), &response)
	require.NoError(t, err)

	assert.Equal(t, config.ExpectedBody, reflect.ValueOf(response).Elem().Interface())
}

func performDeleteTest(t *testing.T, config TestData) {
	delRecorder, err := makeRequest("DELETE", config.Route, nil)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, delRecorder.Code)

	// verify deletion
	getRecorder, err := makeRequest("GET", config.Route, nil)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, getRecorder.Code)
}
