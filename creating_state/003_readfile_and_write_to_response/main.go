package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(write))
	http.Handle("/favicon", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func write(w http.ResponseWriter, r *http.Request) {
	var s string
	method := r.Method
	if method == http.MethodPost {
		// Open file
		f, h, err := r.FormFile("q")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n file:", f, "\n header:", h)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type=file name="q">
	<button>submit</button>
	</form> <br>
	`+s)
}
