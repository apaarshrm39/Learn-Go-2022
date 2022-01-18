package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method of request at foo is", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method of request at Bar is", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The method of request at barred is", r.Method)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form action="/bar" method="POST">
    <button>Submit</button>
</form>
	`)
}
