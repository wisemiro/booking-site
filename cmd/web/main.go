package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/handlers"
	"github.com/wycemiro/booking-site/internal/models"
	"github.com/wycemiro/booking-site/internal/renders"

	"github.com/alexedwards/scs/v2"
)

const port = ":8000"

var app config.AppConfig
var sessions *scs.SessionManager

func main() {
	gob.Register(models.Reservation{})

	//config
	app.InProduction = false //change to true in production, to change secure = true.

	//sessions
	sessions = scs.New()
	sessions.Lifetime = 24 * time.Hour
	sessions.Cookie.Persist = true
	sessions.Cookie.Secure = app.InProduction
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	app.Sessions = sessions

	//templates
	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cant create cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false //if set to true use cache on disk else=false read from file

	renders.CreateTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//server
	fmt.Printf("started server on localhost%s", port)
	serzer := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = serzer.ListenAndServe()
	log.Fatal(err)

}
