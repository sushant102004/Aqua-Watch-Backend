package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName      string             `bson:"firstName" json:"firstName" required:"true"`
	LastName       string             `bson:"lastName" json:"lastName" required:"true"`
	Email          string             `bson:"email" json:"email" required:"true"`
	Location       string             `bson:"location" json:"location" required:"true"`
	FavoritePlace  []string           `bson:"favoritePlace" json:"favoritePlace" required:"true"`
	Language       string             `bson:"language" json:"language" required:"true"`
	ProfilePicture string             `bson:"profilePicture" json:"profilePicture" required:"true"`
}
