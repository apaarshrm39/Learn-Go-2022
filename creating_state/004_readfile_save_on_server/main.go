package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.HandlerFunc(handle))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	var s string
	method := r.Method
	if method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			fmt.Println(err)
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}

		s = string(bs)

		os.WriteFile("files/"+h.Filename, bs, 6666)
	}

	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)

	fmt.Fprint(os.Stdout, string(bs))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type=file name="q">
	<button>submit</button>
	</form> <br>
	`+s)

}
