package main

import (
	"io"
	"net/http"
)

type hotdog int

type cat float32

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1> Doggy <h1>")
}

func (c cat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1> Cat<h1>")
}

func main() {
	var c cat
	var h hotdog
	mux := http.NewServeMux()
	// because there is a forward slash after dog/ it catches also urls like /dog/something
	mux.Handle("/dog/", h)
	// This wil not catch /cat/something
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}
