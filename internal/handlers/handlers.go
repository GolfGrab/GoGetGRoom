package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GolfGrab/GoGetGRoom/internal/config"
	"github.com/GolfGrab/GoGetGRoom/internal/models"
	"github.com/GolfGrab/GoGetGRoom/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})

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
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Contact is the contact page handler render the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is the make-reservation page handler render the make-reservation page with the form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Pink is the pink page handler render the pink page
func (m *Repository) Pink(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "pink.page.tmpl", &models.TemplateData{})
}

// Gray is the gray page handler render the gray page
func (m *Repository) Gray(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "gray.page.tmpl", &models.TemplateData{})
}

//Availabilitty is the search-availabilitty page handler render the search-availabilitty page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

//PostAvailabilitty is the search-availabilitty request handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")

	w.Write([]byte("start date is " + start + " and end date is " + end))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//AvailabilittyJSON is the search-availabilitty request handler that send json response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println("error:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
