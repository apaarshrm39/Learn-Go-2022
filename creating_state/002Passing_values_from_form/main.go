package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		value := r.FormValue("q")
		io.WriteString(rw, `
		<form action="POST">
    <input type="text" name="q">
</form> <br>
		<h1>`+value+"</h1>")
	}))

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
