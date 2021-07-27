package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(mongodbURL string, dbName string) (*mongo.Database, error) {
	log.Println("Starting connection to mongodb at", mongodbURL)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB. Make sure mongodb instance is running.")
		return nil, err
	}

	return client.Database(dbName), nil
}
