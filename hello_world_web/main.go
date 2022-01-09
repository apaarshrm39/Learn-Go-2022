package main

import (
	"errors"
	"fmt"
	"net/http"
)

// use const for portNumber
var portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> This is the response page </h1>")
}

// about is the about page handler
func about(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(1, 4)
	fmt.Fprintf(w, fmt.Sprintf("<h1> This is an about Page, and 1 + 4 is equal to %d </h1", sum))
}

func AddValues(x, y int) int {
	return x + y
}

func divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf(err.Error()))
		return
	}

	fmt.Fprintf(w, fmt.Sprintf(" <h1> %f divided by %f is %f </h1>", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		return 0, errors.New("Cannont divide by zero")
	}
	result := x / y
	return result, nil
}

// main is the main application fuction
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/divide", divide)
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			n, err := fmt.Fprintf(w, "Hello World")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("printed", n, "bytes")
		})
	*/

	fmt.Println(fmt.Sprintf("Starting application at port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
