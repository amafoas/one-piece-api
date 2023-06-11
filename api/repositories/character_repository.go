package repositories

import (
	"context"
	"log"
	"one-piece-api/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCharacterByField(key string, value interface{}) (*models.Character, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	collection := db.Collection("characters")

	character := &models.Character{}
	filter := bson.M{key: value}

	err = collection.FindOne(context.TODO(), filter).Decode(character)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return character, nil
}
