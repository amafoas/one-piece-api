package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CharacterRepository struct {
	*BaseRepository
}

func NewCharacterRepository(db *mongo.Database) *CharacterRepository {
	return &CharacterRepository{
		&BaseRepository{
			Collection: db.Collection("characters"),
		},
	}
}
