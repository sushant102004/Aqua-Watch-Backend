package db

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

const DBNAME = "AquaWatch"

func ConnectToMongo(URI string) *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}
