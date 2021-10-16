package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

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
