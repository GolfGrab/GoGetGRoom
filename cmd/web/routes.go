package main

import (
	"net/http"

	"github.com/GolfGrab/GoGetGRoom/pkg/config"
	"github.com/GolfGrab/GoGetGRoom/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoader)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/gray", handlers.Repo.Gray)
	mux.Get("/pink", handlers.Repo.Pink)
	mux.Get("/check-availability", handlers.Repo.Availability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
