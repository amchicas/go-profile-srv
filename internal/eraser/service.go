package eraser

import (
	"context"

	"github.com/amchicas/go-profile-srv/internal/domain"
	"github.com/amchicas/go-profile-srv/pkg/log"
)

type Service interface {
	DeleteProfileById(ctx context.Context, id uint64) error
}
type service struct {
	repo   domain.Repository
	logger *log.Logger
}

func (s *service) DeleteProfileById(ctx context.Context, id uint64) error {

	err := s.repo.DeleteProfileById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func NewService(repo domain.Repository, logger *log.Logger) Service {

	return &service{
		repo:   repo,
		logger: logger,
	}

}
