package unit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataObject struct {
	ID     string `bson:"_id"`
	Value1 string `bson:"value1"`
	Value2 int    `bson:"value2"`
	Value3 bool   `bson:"value3"`
}

var obj = DataObject{
	ID:     "obj",
	Value1: "string value",
	Value2: 1000,
	Value3: true,
}

type BaseRepositorySuite struct {
	suite.Suite
}

func (suite *BaseRepositorySuite) Test1Create() {
	err := repo.Create(obj)
	require.NoError(suite.T(), err)
}

func (suite *BaseRepositorySuite) Test2FindByID() {
	var findedObj DataObject
	err := repo.FindByID(obj.ID, &findedObj)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), findedObj, obj)
}

func (suite *BaseRepositorySuite) Test3FindNoDocument() {
	var findedObj DataObject
	err := repo.FindByID("no-id", &findedObj)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, mongo.ErrNoDocuments.Error())
}

func (suite *BaseRepositorySuite) Test4Update() {
	update := map[string]interface{}{
		"value1": "updated value",
		"value3": false,
	}

	err := repo.Update(obj.ID, update)
	require.NoError(suite.T(), err)

	var findedObj DataObject
	err = repo.FindByID(obj.ID, &findedObj)
	require.NoError(suite.T(), err)

	updatedObj := obj
	updatedObj.Value1 = "updated value"
	updatedObj.Value3 = false
	require.Equal(suite.T(), updatedObj, findedObj)
}

func (suite *BaseRepositorySuite) Test5Delete() {
	err := repo.Delete(obj.ID)
	require.NoError(suite.T(), err)

	var findedObj DataObject
	err = repo.FindByID(obj.ID, &findedObj)
	require.Error(suite.T(), err)
	require.EqualError(suite.T(), err, mongo.ErrNoDocuments.Error())
}

func TestBaseRepositorySuite(t *testing.T) {
	suite.Run(t, new(BaseRepositorySuite))
}
