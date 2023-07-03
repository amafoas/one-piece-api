package integration

import (
	"fmt"
	"testing"

	"one-piece-api/api/models"

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
	performInsertTest(suite.T(), TestData{
		"/chapter",
		testChapter,
		nil,
	})
}

func (suite *ChapterTestSuite) Test2GetChapterByID() {
	performGetTest(suite.T(), TestData{
		fmt.Sprintf("/chapter/%s", testChapter.ID),
		nil,
		testChapter,
	})
}

func (suite *ChapterTestSuite) Test3UpdateChapter() {
	payload := map[string]interface{}{
		"title": "Updated title",
		"pages": 25,
	}

	updatedChapter := testChapter
	updatedChapter.Title = "Updated title"
	updatedChapter.Pages = 25

	performUpdateTest(suite.T(), TestData{
		fmt.Sprintf("/chapter/%s", testChapter.ID),
		payload,
		updatedChapter,
	})
}

func (suite *ChapterTestSuite) Test4DeleteChapter() {
	performDeleteTest(suite.T(), TestData{
		fmt.Sprintf("/chapter/%s", testChapter.ID),
		nil,
		nil,
	})
}

func TestChapterSuite(t *testing.T) {
	suite.Run(t, new(ChapterTestSuite))
}
