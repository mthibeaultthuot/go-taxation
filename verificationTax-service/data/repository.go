package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	GetAll() []Tax
	GetAllFromUser(username string) []Tax
	Insert(username string) error
	Delete() error
}

func (mongodb *MongoDb) GetAll() []Tax {
	return nil
}

func (mongodb *MongoDb) GetAllFromUser(username string) []Tax {
	coll := Client.Database("Tax").Collection("verification")
	filter := bson.D{{"username", username}}
	cursor, _ := coll.Find(context.TODO(), filter)
	var results []Tax
	_ = cursor.All(context.TODO(), &results)
	return results
}

func (mongodb *MongoDb) Insert(tax *Tax) error {
	coll := mongodb.Client.Database("Tax").Collection("verification")
	_, err := coll.InsertOne(context.TODO(), tax)
	return err
}

func (mongodb *MongoDb) Delete() error {
	return nil
}
