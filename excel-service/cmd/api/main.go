package main

import (
	"log"
	"net/http"
)

type Config struct{}

var (
	ip = "localhost:9001"
)

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    ip,
		Handler: app.route(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	go grpcListen()
}
