package handlers

import (
	"bookings/packs/config"
	"bookings/packs/models"
	"bookings/packs/render"
	"fmt"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	app *config.App
}

// NewRepo creates a new repository
func NewRepo(a *config.App) *Repository {
	return &Repository{
		app: a,
	}
}

// NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got the request for home %s\n", r.Method)
	remoteIp := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote ip", remoteIp)
	render.Template(w, "home.page.html", nil)

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// var td models.TemplateData
	stringMap := make(map[string]string)

	remoteIp := m.app.Session.GetString(r.Context(), "remote ip")
	fmt.Printf("remote ip is %q\n", remoteIp)
	stringMap["remote ip"] = remoteIp
	stringMap["test"] = "this is the string from the handlers"
	fmt.Printf("Got the request for about %s\n", r.Method)
	render.Template(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
