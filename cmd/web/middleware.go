package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func crsfToken(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)
	crsfHandler.SetBaseCookie(
		http.Cookie{
			HttpOnly: true,
			Secure: false,
			Path: "/", 
			SameSite: http.SameSiteLaxMode,

		},
	)
	return crsfHandler
}