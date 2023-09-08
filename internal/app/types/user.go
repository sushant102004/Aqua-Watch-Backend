package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName      string             `bson:"firstName" json:"firstName"`
	LastName       string             `bson:"lastName" json:"lastName"`
	Email          string             `bson:"email" json:"email"`
	Location       string             `bson:"location" json:"location"`
	FavoritePlace  []string           `bson:"favoritePlace" json:"favoritePlace"`
	Language       string             `bson:"language" json:"language"`
	ProfilePicture string             `bson:"profilePicture" json:"profilePicture"`
}
