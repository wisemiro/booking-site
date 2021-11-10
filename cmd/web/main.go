package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/driver"
	"github.com/wycemiro/booking-site/internal/handlers"
	"github.com/wycemiro/booking-site/internal/models"
	"github.com/wycemiro/booking-site/internal/renders"

	"github.com/alexedwards/scs/v2"
)

const port = ":8000"

var app config.AppConfig
var sessions *scs.SessionManager

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close() //close database

	//server
	fmt.Printf("started server on localhost%s", port)
	serzer := http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = serzer.ListenAndServe()
	log.Fatal(err)

}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	//config
	app.InProduction = false //change to true in production, to change secure = true.

	//sessions
	sessions = scs.New()
	sessions.Lifetime = 24 * time.Hour
	sessions.Cookie.Persist = true
	sessions.Cookie.Secure = app.InProduction
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	app.Sessions = sessions

	//connect to database
	log.Println("Connecting to database...🍃")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=api password=50451103aA")
	if err != nil {
		log.Fatal("Can't connect to the database ☠️")
	}
	log.Println("Connected to the database 🎉")
	//templates
	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cant create cache ❌ ", err)
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false //if set to true use cache on disk else=false read from file
	renders.NewRenderer(&app)
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	return db, nil
}
