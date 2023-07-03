package integration

import (
	"fmt"
	"testing"

	"one-piece-api/api/models"

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
	performInsertTest(suite.T(), TestData{
		"/episode",
		testEpisode,
		nil,
	})
}

func (suite *EpisodeTestSuite) Test2GetEpisodeByID() {
	performGetTest(suite.T(), TestData{
		fmt.Sprintf("/episode/%s", testEpisode.ID),
		nil,
		testEpisode,
	})
}

func (suite *EpisodeTestSuite) Test3UpdateEpisode() {
	payload := map[string]interface{}{
		"title":  "Updated title",
		"season": 25,
	}

	updatedEpisode := testEpisode
	updatedEpisode.Title = "Updated title"
	updatedEpisode.Season = 25

	performUpdateTest(suite.T(), TestData{
		fmt.Sprintf("/episode/%s", testEpisode.ID),
		payload,
		updatedEpisode,
	})
}

func (suite *EpisodeTestSuite) Test4DeleteEpisode() {
	performDeleteTest(suite.T(), TestData{
		fmt.Sprintf("/episode/%s", testEpisode.ID),
		nil,
		nil,
	})
}

func TestEpisodeSuite(t *testing.T) {
	suite.Run(t, new(EpisodeTestSuite))
}
