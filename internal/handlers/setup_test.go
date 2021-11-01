package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/models"
	"github.com/wycemiro/booking-site/internal/renders"
)

var app config.AppConfig
var sessions *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler {
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
	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cant create cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = true //if set to true use cache on disk else=false read from file

	repo := NewRepo(&app)
	NewHandlers(repo)
	renders.CreateTemplates(&app)
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer, middleware.Logger, CrsfToken, LoadSession)

	//routes
	mux.Get("/", Repo.Home)

	mux.Get("/about", Repo.About)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/majors-suite", Repo.Major)

	mux.Get("/search-availability", Repo.SearchAvailability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Get("/search-availability-json", Repo.AvailabilityJson)

	mux.Get("/generals-quarters", Repo.General)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	//get static images from static dir
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

//crsfToken generates a token using nosurf
func CrsfToken(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(
		http.Cookie{
			HttpOnly: true,
			Secure:   app.InProduction,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		},
	)
	return crsfHandler
}

//loadSession loads and saves sessions on every request.
func LoadSession(next http.Handler) http.Handler {
	return sessions.LoadAndSave(next)
}

//CreateTemplateCache creates a cache for the templates ?
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
