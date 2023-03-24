package main

import (
	"net/http"

	"github.com/SiddhantGupta07/bookings/pkg/config"
	"github.com/SiddhantGupta07/bookings/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	//returns a http handler
	//Create a multiplexer

	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	//Created my chi router
	//I have a mux
	mux := chi.NewRouter()

	//How to use some library provided middleware with chi
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
