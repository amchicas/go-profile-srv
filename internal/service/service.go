package service

import (
	"context"
	"net/http"
	"time"

	"github.com/amchicas/go-profile-srv/internal/domain"
	"github.com/amchicas/go-profile-srv/pkg/log"
	"github.com/amchicas/go-profile-srv/pkg/pb"
)

type Service struct {
	repo   domain.Repository
	logger *log.Logger
}

func New(repo domain.Repository, logger *log.Logger) *Service {

	return &Service{
		repo:   repo,
		logger: logger,
	}
}
func (s *Service) AddProfile(ctx context.Context, req *pb.AddReq) (*pb.Resp, error) {
	var profile domain.Profile
	_, err := s.repo.GetProfileById(ctx, req.Profile.Id)
	if err == nil {

		return &pb.Resp{
			Status: http.StatusConflict,
			Error:  "Email already  Exist",
		}, nil
	}

	profile.Name = req.Profile.Name
	profile.Lastname = req.Profile.Lastname
	profile.Description = req.Profile.Description
	profile.Votes = req.Profile.Votes
	profile.Students = req.Profile.Students
	profile.Website = req.Profile.Website
	profile.Youtube = req.Profile.Youtube
	profile.Linkedin = req.Profile.Linkedin
	profile.Twitter = req.Profile.Twitter
	profile.Facebook = req.Profile.Facebook
	profile.Created = time.Now().Unix()
	profile.Modified = time.Now().Unix()

	err = s.repo.SaveProfile(ctx, &profile)
	if err != nil {
		s.logger.Error(err.Error())
		return &pb.Resp{
			Status: http.StatusConflict,
			Error:  "Error save  profile",
		}, nil

	}

	return &pb.Resp{
		Status: http.StatusCreated,
		Msg:    "Save Profile",
	}, nil

}

func (s *Service) Update(ctx context.Context, req *pb.UpdateReq) (*pb.Resp, error) {
	profile, err := s.repo.GetProfileById(ctx, req.Profile.Id)
	if err == nil {

		return &pb.Resp{
			Status: http.StatusConflict,
			Error:  "Email already  Exist",
		}, nil
	}

	profile.Name = req.Profile.Name
	profile.Lastname = req.Profile.Lastname
	profile.Description = req.Profile.Description
	profile.Votes = req.Profile.Votes
	profile.Students = req.Profile.Students
	profile.Website = req.Profile.Website
	profile.Youtube = req.Profile.Youtube
	profile.Linkedin = req.Profile.Linkedin
	profile.Twitter = req.Profile.Twitter
	profile.Facebook = req.Profile.Facebook
	profile.Modified = time.Now().Unix()

	err = s.repo.UpdateProfile(ctx, profile, req.Uid)
	if err != nil {
		s.logger.Error(err.Error())
		return &pb.Resp{
			Status: http.StatusConflict,
			Error:  "Error save  profile",
		}, nil

	}

	return &pb.Resp{
		Status: http.StatusCreated,
		Msg:    "Save Profile",
	}, nil

}
func (s *Service) FindProfile(ctx context.Context, req *pb.FindProfileReq) (*pb.FindProfileResp, error) {
	profile, err := s.repo.GetProfileById(ctx, req.Id)
	if err != nil {

		return &pb.FindProfileResp{
			Status: http.StatusInternalServerError,
			Error:  "Email already  Exist",
		}, err
	}

	return &pb.FindProfileResp{
		Status: http.StatusConflict,
		Profile: &pb.Profile{
			Id:          profile.Id,
			Name:        profile.Name,
			Description: profile.Description,
			Votes:       profile.Votes,
			Students:    profile.Students,
			Website:     profile.Website,
			Youtube:     profile.Youtube,
			Linkedin:    profile.Linkedin,
			Twitter:     profile.Twitter,
			Facebook:    profile.Facebook,
			Created:     profile.Created,
			Modified:    profile.Modified,
		},
	}, nil
}
func (s *Service) Remove(ctx context.Context, req *pb.DeleteReq) (*pb.Resp, error) {
	err := s.repo.DeleteProfileById(ctx, req.Uid)
	if err != nil {

		return &pb.Resp{
			Status: http.StatusInternalServerError,
			Error:  "Email already  Exist",
		}, err
	}

	return &pb.Resp{
		Status: http.StatusOK,
		Msg:    "Profile  Deleted",
	}, nil
}
