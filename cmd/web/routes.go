package main

import (
	"github.com/caovanhoang63/bookings/pkg/config"
	"github.com/caovanhoang63/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	//use middleware
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//use Handlers
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	//File server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
