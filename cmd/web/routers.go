package main

import (
	"net/http"

	"github.com/wycemiro/booking-site/internal/config"
	"github.com/wycemiro/booking-site/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer, middleware.Logger, CrsfToken, LoadSession)

	//routes
	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/majors-suite", handlers.Repo.Major)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Get("/search-availability-json", handlers.Repo.AvailabilityJson)

	mux.Get("/generals-quarters", handlers.Repo.General)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	//get static images from static dir
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
