package repositories

type CharacterRepository struct {
	*BaseRepository
}

func NewCharacterRepository() (*CharacterRepository, error) {
	repo, err := NewRepository("characters")
	if err != nil {
		return nil, err
	}

	characterRepo := &CharacterRepository{repo}

	return characterRepo, nil
}
