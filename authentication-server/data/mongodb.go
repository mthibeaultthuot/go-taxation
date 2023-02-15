package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDb interface {
	Init() error
}

var (
	uri    = "mongodb://localhost:27017"
	Client *mongo.Client
)

func Init() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	Client = newClient
	log.Println("Mongodb as started on ", uri)
	return newClient, err
}
