package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Apaar-key", "This is from apaar")
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> Any Code <h2>")
}

func main() {
	mux := http.NewServeMux()
	pointer_mux := &mux
	var d hotdog
	d = 1
	mux.Handle("/poop", d)
	http.ListenAndServe(":8080", *pointer_mux)
}
