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

var testChapter = models.Chapter{
	Title:   "test chapter",
	Volume:  50,
	Chapter: 33,
	Pages:   10,
	Release: "May, 10th",
}

func configChapterRoutesTest(db *mongo.Database) {
	log.Println("Chapter routes test setup")
	routers.ConfigureChapterRoutes(router)

	collection := db.Collection("chapters")

	_, err := collection.InsertOne(context.TODO(), testChapter)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Test chapter inserted in the database.")
}

func TestGetChapterByNumber(t *testing.T) {
	path := fmt.Sprintf("/chapter/%v", testChapter.Chapter)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Incorrect status code")

	expectedBody, err := json.Marshal(testChapter)
	if err != nil {
		t.Error("Error while trying to marshal")
	}
	assert.Equal(t, expectedBody, response.Body.Bytes(), "Wrong Answer Body")
}

func TestGetChapterNotFound(t *testing.T) {
	path := fmt.Sprintf("/chapter/%v", 0)
	req, _ := http.NewRequest("GET", path, nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusNotFound, response.Code, "Incorrect status code")
}
