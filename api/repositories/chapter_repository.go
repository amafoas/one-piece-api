package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ChapterRepository struct {
	*BaseRepository
}

func NewChapterRepository(db *mongo.Database) *ChapterRepository {
	return &ChapterRepository{
		&BaseRepository{
			Collection: db.Collection("chapters"),
		},
	}
}
