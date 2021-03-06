package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("index.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", 42)
	if err != nil {
		log.Fatal(err)
	}
}
