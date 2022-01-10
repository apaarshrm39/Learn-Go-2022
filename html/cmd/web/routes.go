package main

import (
	"net/http"

	"example.com/html/pkg/config"
	"example.com/html/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

// Pat Router
func Route(app *config.AppConfig) http.Handler {

	/*
		mux := pat.New()
		mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

		return mux
	*/

	// Chi Router

	mux := chi.NewRouter()

	// middleware!
	mux.Use(middleware.Recoverer)
	//custom middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux

}
