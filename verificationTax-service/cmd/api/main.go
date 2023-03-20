package main

import (
	"github.com/D3xt3rrrr/verificationTax-service/data"
	"log"
	"net/http"
)

type Config struct {
	Repo data.MongoDb
}

var (
	ip = "localhost:8081"
)

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    ip,
		Handler: app.route(),
	}

	Mongodb := data.MongoDb{
		LogName:    "MongoDb Tax",
		Uri:        "mongodb://localhost:27017",
		Database:   "Tax",
		Collection: "verification",
	}

	client, err := data.Init(Mongodb)
	if err != nil {
		panic(err)
	}
	app.Repo.Client = client

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
