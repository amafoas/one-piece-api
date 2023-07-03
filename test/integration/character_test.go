package integration

import (
	"fmt"
	"one-piece-api/api/models"
	"testing"

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
	performInsertTest(suite.T(), TestData{
		"/character",
		testCharacter,
		nil,
	})
}

func (suite *CharacterTestSuite) Test2GetCharacterByID() {
	performGetTest(suite.T(), TestData{
		fmt.Sprintf("/character/%s", testCharacter.ID),
		nil,
		testCharacter,
	})
}

func (suite *CharacterTestSuite) Test3UpdateCharacter() {
	payload := map[string]interface{}{
		"name":   "Updated name",
		"origin": "Updated origin",
	}

	updatedCharacter := testCharacter
	updatedCharacter.Name = "Updated name"
	updatedCharacter.Origin = "Updated origin"

	performUpdateTest(suite.T(), TestData{
		fmt.Sprintf("/character/%s", testCharacter.ID),
		payload,
		updatedCharacter,
	})
}

func (suite *CharacterTestSuite) Test4DeleteCharacter() {
	performDeleteTest(suite.T(), TestData{
		fmt.Sprintf("/character/%s", testCharacter.ID),
		nil,
		nil,
	})
}

func TestCharacterSuite(t *testing.T) {
	suite.Run(t, new(CharacterTestSuite))
}
