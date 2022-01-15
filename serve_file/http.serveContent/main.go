package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	http.Handle("/dogo", http.HandlerFunc(doggy))

	http.ListenAndServe(":8080", nil)
}

func doggy(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("flux.png")
	if err != nil {
		log.Fatal(err)
	}
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	http.ServeContent(w, r, info.Name(), info.ModTime(), file)
}
