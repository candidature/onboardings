package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/candidature/onboardings/config"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hit the Page")
		next.ServeHTTP(w,r)
	})
}

//Add CSRF protection to every POST request.
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		 HttpOnly: true,
		 Path: "/",
		 Secure: config.GetConfig().InProduction,
		 SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
//Loads and saves session on every request.
func SessionLoad(next http.Handler) http.Handler {
	return config.GetConfig().Scs.LoadAndSave(next)
}