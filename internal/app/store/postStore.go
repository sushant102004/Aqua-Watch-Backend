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
	GetAllPosts(context.Context) ([]types.UserPost, error)
	IncreaseDamageScore(context.Context, string) error
	SearchPostsVIALocation(context.Context, string) ([]types.UserPost, error)
	// This method will only provide data that is relevant to be shown on map with marker.
	GetPostsForMap(context.Context) ([]types.UserPostMap, error)
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

func (s *MongoPostStore) GetAllPosts(ctx context.Context) ([]types.UserPost, error) {
	cursor, err := s.col.Find(ctx, bson.D{})
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

func (s *MongoPostStore) IncreaseDamageScore(ctx context.Context, postID string) error {
	id, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return fmt.Errorf("post id incorrect")
	}

	var post types.UserPost

	res := s.col.FindOne(ctx, bson.M{"_id": id})

	if err := res.Decode(&post); err != nil {
		return fmt.Errorf("error - unable to get post data: %v", err)
	}

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{
		"damageScore": post.DamageScore + 1,
	}}

	_, err = s.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("unable to update damage score: %v", err)
	}

	return nil
}

func (s *MongoPostStore) SearchPostsVIALocation(ctx context.Context, city string) ([]types.UserPost, error) {
	filter := bson.M{"location": bson.M{"$regex": primitive.Regex{Pattern: city, Options: "i"}}}

	res, err := s.col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("unable to search posts for "+city+" %v", err)
	}

	var posts []types.UserPost

	if err := res.All(ctx, &posts); err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return posts, nil
}

func (s *MongoPostStore) GetPostsForMap(ctx context.Context) ([]types.UserPostMap, error) {
	cursor, err := s.col.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var posts []types.UserPostMap
	if err := cursor.All(ctx, &posts); err != nil {
		log.Fatal(err)
	}

	return posts, nil
}
