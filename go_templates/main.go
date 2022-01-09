package main

import (
	"html/template"
	"net/http"
)

const portNumber = ":9090"

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		todos := ToDo{
			User: "Apaar Sharma",
			List: []entry{
				{Name: "Go",
					Done: true,
				},
				{
					Name: "PHP",
					Done: false,
				},
			},
		}
		paths := []string{
			"templates/todo.page.tmpl",
		}
		t, err := template.New("todo.page.tmpl").ParseFiles(paths...)
		err = t.Execute(w, todos)
		if err != nil {
			panic(err)
		}
	})

	// Parse data -- omitted for brevity

	//Providing files as a slice of strings

	http.ListenAndServe(portNumber, nil)

	/*
		paths := []string{
			"templates/todo.page.tmpl",
		}
		name := path.Base(paths[0])

		todos := ToDo{
			User: "Apaar Sharma",
			List: []entry{
				{Name: "Go",
					Done: true,
				},
				{
					Name: "PHP",
					Done: false,
				},
			},
		}
		fmt.Println(name)
		t := template.Must(template.New(name).ParseFiles(paths[0]))
		err := t.Execute(os.Stdout, todos)
		if err != nil {
			panic(err)
		}
	*/

}
