package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	InsertUser() (bool, error)
}

func InsertUser(user User) error {
	coll := Client.Database("Tax").Collection("user")
	_, err := coll.InsertOne(context.TODO(), user)
	return err
}

func FetchUser(username string) (User, error) {
	var user User
	coll := Client.Database("Tax").Collection("user")
	filter := bson.D{{"username", username}}
	err := coll.FindOne(context.TODO(), filter).Decode(&user)
	return user, err
}

func DeleteUser(user User) error {
	coll := Client.Database("Tax").Collection("user")
	_, err := coll.DeleteOne(context.TODO(), user)
	return err
}
