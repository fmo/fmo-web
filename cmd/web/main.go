package main

import (
	"fmt"
	"github.com/fmo/fmo-web/pkg/config"
	"github.com/fmo/fmo-web/pkg/handlers"
	"github.com/fmo/fmo-web/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8777"

func main() {
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app := config.AppConfig{
		TemplateCache: tc,
		UseCache:      false,
	}

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
