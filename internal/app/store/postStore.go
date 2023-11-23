package store

import (
	"context"
	"fmt"
	"log"

	"github.com/sushant102004/aqua-watch-backend/internal/app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStore interface {
	InsertPost(context.Context, *types.UserPost) (string, error)
	GetAllPosts(context.Context, string) ([]types.UserPost, error)
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

func (s *MongoPostStore) GetAllPosts(ctx context.Context, location string) ([]types.UserPost, error) {
	filter := bson.M{"location": bson.M{"$regex": primitive.Regex{Pattern: location, Options: "i"}}}

	cursor, err := s.col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer cursor.Close(ctx)

	var posts []types.UserPost
	if err := cursor.All(ctx, &posts); err != nil {
		log.Fatal(err)
	}

	return posts, nil
}
