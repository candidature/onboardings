package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/candidatures/onboardings/pkg/handlers"
	"net/http"
)

func approutes() http.Handler{
	fmt.Print("Using approutes")
	//mux:= pat.New()

	//mux.Get("/", http.HandlerFunc(handlers.Home))

	//mux.Get("/about", http.HandlerFunc(handlers.About))

	mux:= chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	mux.Get("/",handlers.Home)
	mux.Get("/about",handlers.About)

	return mux
}
