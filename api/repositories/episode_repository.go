package repositories

import (
	"context"
	"log"
	"one-piece-api/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetEpisodeByNumber(num int) (*models.Episode, error) {
	db, err := GetDB()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	collection := db.Collection("episodes")

	episode := &models.Episode{}
	filter := bson.M{"episode": num}

	err = collection.FindOne(context.TODO(), filter).Decode(episode)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return episode, nil
}
