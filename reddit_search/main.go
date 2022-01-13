package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("Reditter.gohtml"))
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		tpl.ExecuteTemplate(rw, "Reditter.gohtml", r.Form)
	}))
	if err != nil {
		log.Fatal(err)
	}
}
