package main

import (
	"html/template"
	"math"
	"os"
)

var tpl *template.Template

var funcmap = template.FuncMap{
	"fdbl":  double,
	"fsqrt": sqroot,
	"fsq":   square,
}

func double(f float32) float32 {
	return 2 * f
}

func sqroot(f float32) float32 {
	return float32(math.Sqrt(float64(f)))
}

func square(f float32) float32 {
	return float32(math.Pow(float64(f), 2))
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcmap).ParseFiles("math.gohtml"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout, "math.gohtml", float32(64.0))
}
