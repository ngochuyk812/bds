package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient(ctx context.Context, connection string) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	println("Connected to MongoDB")
	return client
}
