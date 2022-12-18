package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// adding state to rememeber the session
// SessionLoad loads and savs the session on every request
// nb on every request, not every page load: it can be fired more than once if the pages requests multiple files!
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
