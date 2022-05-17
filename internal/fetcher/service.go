package fetcher

import (
	"context"

	"github.com/amchicas/go-profile-srv/internal/domain"
	"github.com/amchicas/go-profile-srv/pkg/log"
)

type Service interface {
	GetProfile(ctx context.Context, id uint64) (*domain.Profile, error)
}
type service struct {
	repo   domain.Repository
	logger *log.Logger
}

func (s *service) GetProfile(ctx context.Context, id uint64) (*domain.Profile, error) {

	profile, err := s.repo.GetProfileById(ctx, id)
	if err != nil {

		return &domain.Profile{}, err
	}

	return profile, nil
}
func NewService(repo domain.Repository, logger *log.Logger) Service {

	return &service{
		repo:   repo,
		logger: logger,
	}

}
