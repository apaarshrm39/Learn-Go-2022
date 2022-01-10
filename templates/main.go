package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t, err := template.ParseFiles("./tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("index.html")
	defer file.Close()

	//t.Execute(file, nil)
	t, err = t.ParseFiles("index.html")
	//t.Execute(os.Stdout, nil)

	// If my template contains more than one template then I can use execute Templates

	t.ExecuteTemplate(os.Stdout, "index.html", nil)
	//fmt.Println(t)

	tpl, err := template.ParseGlob("html/*")
	tpl.ExecuteTemplate(os.Stdout, "two.html", nil)

}
