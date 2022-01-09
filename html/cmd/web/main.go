package main

import (
	"fmt"
	"net/http"

	"example.com/html/pkg/handlers"
)

// use const for portNumber
const portNumber = ":8080"

// main is the main application fuction
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application at port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
