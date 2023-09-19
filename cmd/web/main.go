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
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)

}
