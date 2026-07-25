package favorite

import "context"

type Service struct {
	repository Repository
}

func NewService(repository Repository) UseCase {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, favorite *Favorite) (*Favorite, error) {
	return s.repository.Create(ctx, favorite)
}
