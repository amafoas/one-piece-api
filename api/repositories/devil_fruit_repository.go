package repositories

type DevilFruitRepository struct {
	*BaseRepository
}

func NewDevilFruitRepository() (*DevilFruitRepository, error) {
	repo, err := NewRepository("devil_fruits")
	if err != nil {
		return nil, err
	}

	devilFruitRepo := &DevilFruitRepository{repo}

	return devilFruitRepo, nil
}
