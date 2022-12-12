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

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cahce")
	}

	app.TemplateCache = tc

	// In development mode set to false so any change to a template will be reflected in the browser without having to restart the server
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Server up and running on port", port)
	http.ListenAndServe(port, nil)
}
