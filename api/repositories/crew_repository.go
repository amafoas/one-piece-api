package repositories

import (
	"context"
	"log"
	"one-piece-api/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCrewByID(id string) (*models.Crew, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	collection := db.Collection("crews")

	crew := &models.Crew{}
	filter := bson.M{"_id": id}

	err = collection.FindOne(context.TODO(), filter).Decode(crew)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return crew, nil
}
