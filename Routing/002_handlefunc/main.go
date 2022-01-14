package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cat", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "<h1> HandlerFunc <h1>")
	})
	http.Handle("/dog", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OOga Boooga")
	}))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
