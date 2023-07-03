package repositories

type CrewRepository struct {
	*BaseRepository
}

func NewCrewRepository() (*CrewRepository, error) {
	repo, err := NewRepository("crews")
	if err != nil {
		return nil, err
	}

	crewRepo := &CrewRepository{repo}

	return crewRepo, nil
}
