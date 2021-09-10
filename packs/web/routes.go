package main

import (
	"bookings/packs/config"
	"bookings/packs/handlers"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.App) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	staticDir := http.Dir("./static/")
	fmt.Println(staticDir)
	fmt.Println(filepath.Glob(string(staticDir)))
	fileServer := http.FileServer(http.Dir("./packs/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
