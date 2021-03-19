package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDB(mongoURI, collection string) *MongoDB {
	db := ConnectMongoDB(mongoURI)
	coll := db.Collection(collection)
	return &MongoDB{
		Database:   db,
		Collection: coll,
	}
}

func ConnectMongoDB(uri string) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Database Connection Error ", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Database Connection Error ", err)
	}

	return client.Database("sigma")
}
