package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hakansenn/bookings/pkg/config"
	"github.com/hakansenn/bookings/pkg/handlers"
	"github.com/hakansenn/bookings/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache (main)")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, routes(&app))

}
