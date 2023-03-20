package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDb struct {
	Client     *mongo.Client
	LogName    string
	Uri        string
	Database   string
	Collection string
}

var (
	Client *mongo.Client
)

func Init(mongodb MongoDb) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb.Uri))
	Client = newClient
	log.Println("Mongodb as started on ", mongodb.Uri)
	return newClient, err
}
