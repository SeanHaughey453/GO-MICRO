package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	//"github.com/gorilla/mux"
)

func (app *Config) routes() http.Handler {
	mux:= chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Autorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	return mux
}