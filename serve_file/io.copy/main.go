package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		f, err := os.Open("flux.png")
		if err != nil {
			io.WriteString(rw, "404 Not Found")
		}
		io.Copy(rw, f)
	})
	http.ListenAndServe(":8080", nil)
}
