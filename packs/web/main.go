package main

import (
	"bookings/packs/config"
	"bookings/packs/handlers"
	"bookings/packs/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const myPort = ":8080"

var app config.App
var session *scs.SessionManager

func main() {

	devEnv := true

	// Modify in production
	inProd := false

	// Set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd

	// Setup system configeration
	var err error
	app.InProd = inProd
	app.Session = session

	// Setup render
	app.TemplateCache, err = render.CreateTmplCache()
	if err != nil {
		log.Fatal("error while parsing templates", err)
	}
	render.Config(&app)

	// Setup handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	if devEnv {
		app.UseCache = false
	}

	fmt.Println("\nsetting routes and starting server")

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.ListenAndServe(myPort, nil)

	srv := &http.Server{
		Addr:    myPort,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
