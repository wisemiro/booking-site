package main

import (
	"fmt"
	"learn/pkgs/config"
	"learn/pkgs/handlers"
	"learn/pkgs/renders"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	//config
	var app config.AppConfig

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cant create cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false //if set to true use cache on disk else=false read from file

	renders.CreateTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//urls
	http.HandleFunc("/", handlers.Repo.Home)

	//server
	fmt.Printf("starting server on %s", port)
	http.ListenAndServe(port, nil)
}
