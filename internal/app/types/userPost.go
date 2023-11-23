package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserPost struct {
	UserID      primitive.ObjectID `bson:"user" json:"user"`
	Date        string             `bson:"date" json:"date"`
	Time        string             `bson:"time" json:"time"`
	ImageURL    string             `bson:"imageUrl" json:"imageUrl"`
	Description string             `bson:"description" json:"description"`
	DamageScore int                `bson:"damageScore" json:"damageScore"`
	Coordinates []float64          `bson:"coordinates" json:"coordinates"`
	Location    string             `bson:"location" json:"location"`
}
