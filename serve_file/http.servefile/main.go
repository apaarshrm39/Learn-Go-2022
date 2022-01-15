package main

import (
	"net/http"
)

func main() {

	http.Handle("/dogo", http.HandlerFunc(doggy))

	http.ListenAndServe(":8080", nil)
}

func doggy(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "flux.png")
}
