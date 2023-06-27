package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CrewRepository struct {
	*BaseRepository
}

func NewCrewRepository(db *mongo.Database) *CrewRepository {
	return &CrewRepository{
		&BaseRepository{
			Collection: db.Collection("crews"),
		},
	}
}
