package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		value := r.FormValue("q")
		io.WriteString(rw, "Do my search"+value)
	}))
	http.ListenAndServe(":8080", nil)
}
