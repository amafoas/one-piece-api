package repositories

type EpisodeRepository struct {
	*BaseRepository
}

func NewEpisodeRepository() (*EpisodeRepository, error) {
	repo, err := NewRepository("episodes")
	if err != nil {
		return nil, err
	}

	episodeRepo := &EpisodeRepository{repo}

	return episodeRepo, nil
}
