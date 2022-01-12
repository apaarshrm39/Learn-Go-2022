package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	err := http.ListenAndServe(":8080", d)
	if err != nil {
		log.Fatalln(err)
	}
}
