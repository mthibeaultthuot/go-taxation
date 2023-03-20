package main

import (
	"github.com/D3xt3rrrr/go-cloud/authentication-server/data"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Config struct {
	Client *mongo.Client
}

var (
	ip = "0.0.0.0:9001"
)

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    ip,
		Handler: app.route(),
	}

	log.Println("rest api server start on : ", ip)

	mongodb := data.MongoDb{
		"MongoDb Tax",
		"mongodb://root:9144tbbw@mongodb:27017",
		"Tax",
		"user",
	}

	init, err := data.Init(mongodb)
	app.Client = init

	if err != nil {
		log.Panic(err)
	}

	go grpcListen()

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
