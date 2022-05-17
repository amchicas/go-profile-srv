package domain

import (
	"context"
	"time"
)

type Profile struct {
	Uid         uint64 `json:"uid"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Votes       uint64 `json:"votes"`
	Students    uint64 `json:"students"`
	Website     string `json:"website"`
	Youtube     string `json:"youtube"`
	Linkedin    string `json:"linkedin"`
	Twitter     string `json:"twitter"`
	Facebook    string `json:"facebook"`
	Created     int64  `json:"created"`
	Modified    int64  `json:"modified"`
}

func NewProfile(name, lastname, title, description, website, youtube, linkedin, twitter, facebook string, votes, students uint64) *Profile {

	return &Profile{
		Name:        name,
		Lastname:    lastname,
		Description: description,
		Votes:       votes,
		Students:    students,
		Website:     website,
		Youtube:     youtube,
		Linkedin:    linkedin,
		Twitter:     twitter,
		Facebook:    facebook,
		Created:     time.Now().Unix(),
		Modified:    time.Now().Unix(),
	}

}

type Repository interface {
	SaveProfile(ctx context.Context, profile *Profile) error
	UpdateProfile(ctx context.Context, profile *Profile, id uint64) error
	GetProfileById(ctx context.Context, id uint64) (*Profile, error)
	DeleteProfileById(ctx context.Context, id uint64) error
}
