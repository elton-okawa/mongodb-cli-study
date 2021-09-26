package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// TODO get host and credentials via config
const HOST string = "mongodb://root:root@localhost:27017"

var instance *mongo.Client

func CreateClient(ctx context.Context) *mongo.Client {
	if instance != nil {
		return instance
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(HOST))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	instance = client
	return instance
}
