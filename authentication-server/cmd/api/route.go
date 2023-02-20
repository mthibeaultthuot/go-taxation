package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func (app *Config) route() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

	})
	// Public route
	r.Use(middleware.Heartbeat("/ping"))
	r.Post("/v1/auth/login", app.Login)
	r.Post("/v1/auth/registration", app.Registration)
	r.Post("/v1/auth/delete", app.Delete)
	r.Get("/v1/auth/refresh/token", app.Refresh)
	r.Get("/v1/auth/token", app.Token)
	// TODO : séparer l'api privér et public

	return r
}
