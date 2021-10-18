package main

import (
	"github.com/wycemiro/booking-site/pkgs/config"
	"github.com/wycemiro/booking-site/pkgs/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer, middleware.Logger, CrsfToken, LoadSession)
	mux.Get("/", handlers.Repo.Home)

	return mux
}