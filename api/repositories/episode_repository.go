package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type EpisodeRepository struct {
	*BaseRepository
}

func NewEpisodeRepository(db *mongo.Database) *EpisodeRepository {
	return &EpisodeRepository{
		&BaseRepository{
			Collection: db.Collection("episodes"),
		},
	}
}
