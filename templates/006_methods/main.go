package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func (p person) Something() int {
	return 7
}

func (p person) Double(x int) int {
	return 2 * x
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p := person{
		Name: "Donnkey",
		Age:  9,
	}

	tpl.Execute(os.Stdout, p)
}
