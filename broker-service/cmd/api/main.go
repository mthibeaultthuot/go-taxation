package main

import (
	"log"
	"net/http"
)

const (
	ip   string = "localhost:8181"
	port string = "8080"
)

type Config struct{}

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
}
