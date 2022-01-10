package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"example.com/html/pkg/config"
	"example.com/html/pkg/models"
)

// funcmap is a map of function that can be used in a template, for example manipulate date etc
var functions = template.FuncMap{}

var app *config.AppConfig

// newTemplates sets the Config for the template Package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Default Data that i want on all pages.
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate function used to render templates

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var templateCache map[string]*template.Template
	if app.UseCache {
		// get the template Cache from Appconfig
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	t := templateCache[tmpl]
	/*
		if !ok {
			log.Fatal("Could not get templates from template Cache")
		}*/

	// put the value of templates from memory to buffer
	buf := newFunction()
	// Add default Data to TD
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser:", err)
		return
	}

	//parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	/*
		err = templateCache[tmpl].Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template:", err)
			return
		}
	*/
}

func newFunction() *bytes.Buffer {
	buf := new(bytes.Buffer)
	return buf
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	// create a map for ready to use templates
	myCache := map[string]*template.Template{}

	// get a glob of all the files in the give path
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Iterate over all these files.
	for _, page := range pages {
		name := filepath.Base(page) //extract just the name

		fmt.Println("Page is currently", page)
		// create a templateset

		// Creates a newHtml template with the given name.
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// check if a layout is defined for the said template ?

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// if a layout is found
		if len(matches) > 0 {
			// craetes resulting template using layouts
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
