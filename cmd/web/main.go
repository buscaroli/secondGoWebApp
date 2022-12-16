package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/buscaroli/secondGoWebApp/pkg/config"
	"github.com/buscaroli/secondGoWebApp/pkg/handlers"
	"github.com/buscaroli/secondGoWebApp/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	// create the templates for every page and save them into a cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// the cached templates are saven in the global config file
	app.TemplateCache = tc

	// In development mode set to false so any change to a template will be reflected in the browser without having to restart the server
	app.UseCache = true

	// the next two lines allows us to access the config from within the handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// the next line allows us to access the config from within the render package
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.ListenAndServe(port, nil)

	fmt.Println("Server up and running on port", port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
