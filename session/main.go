package main

import (
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", temp)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func temp(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("id")
	if err != nil {
		id := uuid.Must(uuid.NewV4())
		c := &http.Cookie{
			Name:  "id",
			Value: id,
		}
		http.SetCookie(w, c)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", cookie)
}
