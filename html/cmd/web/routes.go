package main

import (
	"net/http"

	"example.com/html/pkg/config"
	"example.com/html/pkg/handlers"
	"github.com/bmizerany/pat"
)

func Route(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux

}
