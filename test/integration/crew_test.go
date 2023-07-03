package integration

import (
	"fmt"
	"testing"

	"one-piece-api/api/models"

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
	performInsertTest(suite.T(), TestData{
		"/crew",
		testCrew,
		nil,
	})
}

func (suite *CrewTestSuite) Test2GetCrewByID() {
	performGetTest(suite.T(), TestData{
		fmt.Sprintf("/crew/%s", testCrew.ID),
		nil,
		testCrew,
	})
}

func (suite *CrewTestSuite) Test3UpdateCrew() {
	payload := map[string]interface{}{
		"name":    "Updated name",
		"captain": "Updated captain",
	}

	updatedCrew := testCrew
	updatedCrew.Name = "Updated name"
	updatedCrew.Captain = "Updated captain"

	performUpdateTest(suite.T(), TestData{
		fmt.Sprintf("/crew/%s", testCrew.ID),
		payload,
		updatedCrew,
	})
}

func (suite *CrewTestSuite) Test4DeleteCrew() {
	performDeleteTest(suite.T(), TestData{
		fmt.Sprintf("/crew/%v", testCrew.ID),
		nil,
		nil,
	})
}

func TestCrewSuite(t *testing.T) {
	suite.Run(t, new(CrewTestSuite))
}
