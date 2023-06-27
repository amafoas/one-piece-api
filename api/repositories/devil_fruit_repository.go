package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DevilFruitRepository struct {
	*BaseRepository
}

func NewDevilFruitRepository(db *mongo.Database) *DevilFruitRepository {
	return &DevilFruitRepository{
		&BaseRepository{
			Collection: db.Collection("devil_fruits"),
		},
	}
}
