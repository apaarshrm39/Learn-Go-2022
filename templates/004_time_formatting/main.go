package main

import (
	"html/template"
	"os"
	"strings"
	"time"
)

var funcmap = template.FuncMap{
	"ft": formatTime,
}
var tpl *template.Template

func formatTime(t time.Time) string {
	st := t.String()
	st = strings.TrimSpace(st)
	st = st[:10]
	return st

}

func init() {
	tpl = template.Must(template.New("").Funcs(funcmap).ParseFiles("time.gohtml"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout, "time.gohtml", time.Now())
}
