package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Caching templates: tc is a map that stores the string:template of every page that has already been visited.
// eg []
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check if template already in the map
	_, inMap := tc[t]

	if !inMap {
		// create the template and add it to the map
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// read the template from the map
		log.Println("Using cached page")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to the map (our cache)
	tc[t] = tmpl
	return nil
}
