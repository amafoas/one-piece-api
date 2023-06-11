package repositories

import (
	"context"
	"log"
	"one-piece-api/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetChapterByNumber(num int) (*models.Chapter, error) {
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	collection := db.Collection("chapters")

	chapter := &models.Chapter{}
	filter := bson.M{"chapter": num}

	err = collection.FindOne(context.TODO(), filter).Decode(chapter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return chapter, nil
}
