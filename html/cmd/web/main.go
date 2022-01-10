package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/html/pkg/config"
	"example.com/html/pkg/handlers"
	"example.com/html/pkg/render"
	"github.com/alexedwards/scs/v2"
)

// use const for portNumber
const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is the main application fuction
func main() {

	// initializing a session

	// change this to true when in Production

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 12 * time.Hour              // ttl of cookie
	session.Cookie.Persist = true                  // should the cookie persist after the user closes the browser
	session.Cookie.SameSite = http.SameSiteLaxMode //default
	session.Cookie.Secure = app.InProduction       //set as true in Prod

	app.Session = session

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
