package grpc

import (
	"context"
	"net/http"
	"time"

	"github.com/amchicas/go-profile-srv/internal/adder"
	"github.com/amchicas/go-profile-srv/internal/eraser"
	"github.com/amchicas/go-profile-srv/internal/fetcher"
	"github.com/amchicas/go-profile-srv/internal/modifier"
	"github.com/amchicas/go-profile-srv/pkg/log"
	"github.com/amchicas/go-profile-srv/pkg/pb"
)

type profileHandler struct {
	aS     adder.Service
	fS     fetcher.Service
	mS     modifier.Service
	eS     eraser.Service
	logger *log.Logger
}

func NewHandler(adderService adder.Service, modifieService modifier.Service, fetcherService fetcher.Service, eraserService eraser.Service, logger *log.Logger) pb.ProfileServiceServer {

	return &profileHandler{
		aS:     adderService,
		mS:     modifieService,
		fS:     fetcherService,
		eS:     eraserService,
		logger: logger,
	}
}
func (s *profileHandler) AddProfile(ctx context.Context, req *pb.AddReq) (*pb.Resp, error) {
	_, err := s.fS.GetProfile(ctx, req.Profile.Id)
	if err == nil {

		return &pb.Resp{
			Status: http.StatusConflict,
			Error:  "Profile already  Exist",
		}, err
	}

	err = s.aS.AddProfile(ctx,
		req.Profile.Name,
		req.Profile.Lastname,
		req.Profile.Title,
		req.Profile.Description,
		req.Profile.Website,
		req.Profile.Youtube,
		req.Profile.Linkedin,
		req.Profile.Twitter,
		req.Profile.Facebook,
		req.Profile.Votes,
		req.Profile.Students,
		req.Profile.Id,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return &pb.Resp{
			Status: http.StatusConflict,
			Error:  "Error saving profile",
		}, err
	}

	return &pb.Resp{
		Status: http.StatusCreated,
	}, nil

}

func (s *profileHandler) Update(ctx context.Context, req *pb.UpdateReq) (*pb.Resp, error) {
	profile, err := s.fS.GetProfile(ctx, req.Profile.Id)
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

	err = s.mS.UpdateProfile(ctx,
		req.Profile.Name,
		req.Profile.Lastname,
		req.Profile.Title,
		req.Profile.Description,
		req.Profile.Website,
		req.Profile.Youtube,
		req.Profile.Linkedin,
		req.Profile.Twitter,
		req.Profile.Facebook,
		req.Profile.Votes,
		req.Profile.Students,
		req.Profile.Id,
	)
	if err != nil {
		s.logger.Error(err.Error())
		return &pb.Resp{
			Status: http.StatusConflict,
			Error:  "Error saving  profile",
		}, err

	}

	return &pb.Resp{
		Status: http.StatusOK,
		Msg:    "Profile saved",
	}, nil

}
func (s *profileHandler) FindProfile(ctx context.Context, req *pb.FindProfileReq) (*pb.FindProfileResp, error) {
	profile, err := s.fS.GetProfile(ctx, req.Uid)
	if err != nil {

		return &pb.FindProfileResp{
			Status: http.StatusInternalServerError,
			Error:  "Email already  exist",
		}, err
	}

	return &pb.FindProfileResp{
		Status: http.StatusOK,
		Profile: &pb.Profile{
			Id:          profile.Uid,
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
func (s *profileHandler) Remove(ctx context.Context, req *pb.DeleteReq) (*pb.Resp, error) {
	err := s.eS.DeleteProfileById(ctx, req.Uid)
	if err != nil {

		return &pb.Resp{
			Status: http.StatusInternalServerError,
			Error:  "Could not be  deleted ",
		}, err
	}

	return &pb.Resp{
		Status: http.StatusOK,
	}, nil
}
