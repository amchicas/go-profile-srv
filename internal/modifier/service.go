package modifier

import (
	"context"
	"time"

	"github.com/amchicas/go-profile-srv/internal/domain"
	"github.com/amchicas/go-profile-srv/pkg/log"
)

type Service interface {
	UpdateProfile(ctx context.Context, name, lastname, title, description, website, youtube, linkedin, twitter, facebook string, votes, students, id uint64) error
}

type service struct {
	repo   domain.Repository
	logger *log.Logger
}

func (s *service) UpdateProfile(ctx context.Context, name, lastname, title, description, website, youtube, linkedin, twitter, facebook string, votes, students, id uint64) error {
	profile, err := s.repo.GetProfileById(ctx, id)
	if err != nil {
		return err

	}

	profile.Name = name
	profile.Lastname = lastname
	profile.Description = description
	profile.Votes = votes
	profile.Students = students
	profile.Website = website
	profile.Youtube = youtube
	profile.Linkedin = linkedin
	profile.Twitter = twitter
	profile.Facebook = facebook
	profile.Modified = time.Now().Unix()

	err = s.repo.UpdateProfile(ctx, profile, id)
	if err != nil {
		s.logger.Error(err.Error())
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
