package adder

import (
	"context"

	"github.com/amchicas/go-profile-srv/internal/domain"
	"github.com/amchicas/go-profile-srv/pkg/log"
)

type Service interface {
	AddProfile(ctx context.Context, name, lastname, title, description, website, youtube, linkedin, twitter, facebook string, votes, students, id uint64) error
}

type service struct {
	repo   domain.Repository
	logger *log.Logger
}

func (s *service) AddProfile(ctx context.Context,
	name, lastname, title, description, website, youtube, linkedin, twitter, facebook string,
	votes, students, id uint64,
) error {
	profile := domain.NewProfile(name, lastname, title, description, website, youtube, linkedin, twitter, facebook, votes, students)
	err := s.repo.SaveProfile(ctx, profile)
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
