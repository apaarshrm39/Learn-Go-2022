package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	slice := []string{"hello", "my", "friend"}

	tpl.ExecuteTemplate(os.Stdout, "map.html", slice)
}
