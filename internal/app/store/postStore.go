package store

import (
	"context"
	"fmt"

	"github.com/sushant102004/aqua-watch-backend/internal/app/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStore interface {
	InsertPost(context.Context, *types.UserPost) (string, error)
}

type MongoPostStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongoPostStore(client *mongo.Client) *MongoPostStore {
	col := client.Database("AquaWatch").Collection("posts")

	return &MongoPostStore{
		client: client,
		col:    col,
	}
}

func (s *MongoPostStore) InsertPost(ctx context.Context, post *types.UserPost) (string, error) {
	_, err := s.col.InsertOne(ctx, post)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return "post added successfully", nil
}
