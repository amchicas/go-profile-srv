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
func (r *repository) DeleteProfileById(ctx context.Context, uid int64) error {
	c := r.db.Collection("profile")
	_, err := c.DeleteOne(ctx, bson.M{"id": uid})
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) GetProfileById(ctx context.Context, id int64) (*domain.Profile, error) {
	var profile domain.Profile
	c := r.db.Collection("profile")
	err := c.FindOne(ctx, bson.M{"id": id}).Decode(&profile)
	if err != nil {
		return &domain.Profile{}, err
	}
	return &profile, nil
}
func (r *repository) UpdateProfile(ctx context.Context, profile *domain.Profile, id int64) error {
	c := r.db.Collection("posts")
	result, err := c.UpdateOne(ctx, bson.M{"id": id}, bson.D{
		{"$set", bson.D{
			{"title", profile.Title},
			{"description", profile.Description},
			{"update_at", profile.Modified},
		}},
	})
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}
