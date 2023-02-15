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
	ip = "localhost:9000"
)

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    ip,
		Handler: app.route(),
	}

	init, err := data.Init()
	app.Client = init

	if err != nil {
		log.Panic(err)
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
