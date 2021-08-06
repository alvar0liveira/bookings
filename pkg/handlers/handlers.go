package handlers

import (
	"net/http"

	"github.com/alvaroliveira20/bookings/cmd/models"
	"github.com/alvaroliveira20/bookings/pkg/config"
	"github.com/alvaroliveira20/bookings/pkg/render"
)

// TemplateData holds data sent from handlers to templates

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo Creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the Repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
