package repositories

import (
	"context"
	"log"
	"one-piece-api/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetDevilFruitByID(id string) (*models.DevilFruit, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	collection := db.Collection("devil_fruits")

	fruit := &models.DevilFruit{}
	filter := bson.M{"_id": id}

	err = collection.FindOne(context.TODO(), filter).Decode(fruit)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return fruit, nil
}
