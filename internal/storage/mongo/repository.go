package mongo

import (
	"context"
	"fmt"

	"github.com/amchicas/go-profile-srv/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) domain.Repository {
	return &repository{db: db}
}

func (r *repository) SaveProfile(ctx context.Context, profile *domain.Profile) error {
	c := r.db.Collection("profile")
	_, err := c.InsertOne(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) DeleteProfileById(ctx context.Context, uid uint64) error {
	c := r.db.Collection("profile")
	_, err := c.DeleteOne(ctx, bson.M{"uid": uid})
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) GetProfileById(ctx context.Context, id uint64) (*domain.Profile, error) {
	var profile domain.Profile
	c := r.db.Collection("profile")
	err := c.FindOne(ctx, bson.M{"uid": id}).Decode(&profile)
	if err != nil {
		return &domain.Profile{}, err
	}
	return &profile, nil
}
func (r *repository) UpdateProfile(ctx context.Context, profile *domain.Profile, id uint64) error {
	c := r.db.Collection("profile")
	fmt.Println(profile)
	_, err := c.UpdateOne(ctx, bson.M{"uid": id},
		bson.M{
			"$set": bson.M{
				"name":        profile.Name,
				"description": profile.Description,
				"votes":       profile.Votes,
				"students":    profile.Students,
				"website":     profile.Website,
				"youtube":     profile.Youtube,
				"linkedin":    profile.Linkedin,
				"twitter":     profile.Twitter,
				"facebook":    profile.Facebook,
				"created":     profile.Created,
				"modified":    profile.Modified,
			}},
	)
	if err != nil {
		return err
	}

	return nil
}
