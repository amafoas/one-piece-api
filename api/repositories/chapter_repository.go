package repositories

type ChapterRepository struct {
	*BaseRepository
}

func NewChapterRepository() (*ChapterRepository, error) {
	repo, err := NewRepository("chapters")
	if err != nil {
		return nil, err
	}

	chapterRepo := &ChapterRepository{repo}

	return chapterRepo, nil
}
