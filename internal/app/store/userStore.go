package store

import (
	"context"
	"fmt"

	"github.com/sushant102004/aqua-watch-backend/internal/app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	Login(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	col    *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	col := client.Database("AquaWatch").Collection("citizen")

	return &MongoUserStore{
		client: client,
		col:    col,
	}
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	resp := s.col.FindOne(ctx, user)

	if resp.Err() == mongo.ErrNoDocuments {
		resp, err := s.col.InsertOne(ctx, user)
		if err != nil {
			return nil, err
		}
		user.ID = resp.InsertedID.(primitive.ObjectID)
	} else {
		return nil, fmt.Errorf("user already exists")
	}

	return user, nil
}

func (s *MongoUserStore) Login(ctx context.Context, email string) (*types.User, error) {
	var user *types.User
	resp := s.col.FindOne(ctx, bson.M{"email": email})
	if resp.Err() == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("account not found. please create one")
	}

	resp.Decode(&user)
	return user, nil
}
