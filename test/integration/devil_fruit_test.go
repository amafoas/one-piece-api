package integration

import (
	"fmt"
	"testing"

	"one-piece-api/api/models"

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
	performInsertTest(suite.T(), TestData{
		"/devil-fruit",
		testDevilFruit,
		nil,
	})
}

func (suite *DevilFruitTestSuite) Test2GetDevilFruitByID() {
	performGetTest(suite.T(), TestData{
		fmt.Sprintf("/devil-fruit/%v", testDevilFruit.ID),
		nil,
		testDevilFruit,
	})
}

func (suite *DevilFruitTestSuite) Test3UpdateDevilFruit() {
	payload := map[string]interface{}{
		"name": "Updated name",
		"type": "Updated type",
	}

	updatedDevilFruit := testDevilFruit
	updatedDevilFruit.Name = "Updated name"
	updatedDevilFruit.Type = "Updated type"

	performUpdateTest(suite.T(), TestData{
		fmt.Sprintf("/devil-fruit/%v", testDevilFruit.ID),
		payload,
		updatedDevilFruit,
	})
}

func (suite *DevilFruitTestSuite) Test4DeleteDevilFruit() {
	performDeleteTest(suite.T(), TestData{
		fmt.Sprintf("/devil-fruit/%v", testDevilFruit.ID),
		nil,
		nil,
	})
}

func TestDevilFruitSuite(t *testing.T) {
	suite.Run(t, new(DevilFruitTestSuite))
}
