package store

import (
	"context"
	"fmt"

	"github.com/sushant102004/aqua-watch-backend/internal/app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NGOStore interface {
	SignUp(context.Context, types.NGO) error
	Login(context.Context, string) (*types.NGO, error)
}

type MongoNGOStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongoNGOStore(client *mongo.Client) *MongoNGOStore {
	col := client.Database("AquaWatch").Collection("ngo")

	return &MongoNGOStore{
		client: client,
		col:    col,
	}
}

func (s *MongoNGOStore) SignUp(ctx context.Context, data types.NGO) error {
	_, err := s.col.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoNGOStore) Login(ctx context.Context, email string) (*types.NGO, error) {
	var ngo *types.NGO

	resp := s.col.FindOne(ctx, bson.M{"email": email})
	if resp.Err() == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("account not found. please create one")
	}

	resp.Decode(&ngo)
	return ngo, nil
}
