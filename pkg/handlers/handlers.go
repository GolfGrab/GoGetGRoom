package handlers

import (
	"net/http"

	"github.com/GolfGrab/GoGetGRoom/pkg/config"
	"github.com/GolfGrab/GoGetGRoom/pkg/models"
	"github.com/GolfGrab/GoGetGRoom/pkg/render"
)

//Repo is the repository for the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo is a helper function to create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers is a helper function to sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler render the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// About is the about page handler render the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World from the string map"
	stringMap["name"] = "Golf Grab"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp
	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Contact is the contact page handler render the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is the make-reservation page handler render the make-reservation page with the form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Pink is the pink page handler render the pink page
func (m *Repository) Pink(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "pink.page.tmpl", &models.TemplateData{})
}

// Gray is the gray page handler render the gray page
func (m *Repository) Gray(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "gray.page.tmpl", &models.TemplateData{})
}

//Availabilitty is the search-availabilitty page handler render the search-availabilitty page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}
