package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/html/pkg/config"
	"example.com/html/pkg/handlers"
	"example.com/html/pkg/render"
)

// use const for portNumber
const portNumber = ":8080"

// main is the main application fuction
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Could not Create Template Cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Route(&app),
	}
	//http.HandleFunc("/", repo.Home)
	//http.HandleFunc("/about", repo.About)

	fmt.Println(fmt.Sprintf("Starting application at port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
