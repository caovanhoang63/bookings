package handlers

import (
	"fmt"
	"github.com/caovanhoang63/bookings/pkg/config"
	"github.com/caovanhoang63/bookings/pkg/models"
	"github.com/caovanhoang63/bookings/pkg/render"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Declare handler functions

//ALL HANDLERS MUST HAVE THIS FORMAT
//func (m *Repository) HandlerName(w http.ResponseWriter, r *http.Request)

//Handler for GET request

// Home is the home page handler for GET request
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})

}

// About is the about page handler for GET request
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

// Contact is the contact page handler for GET request
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})

}

// Generals is the generals page handler for GET request
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})

}

// Major is the major page handler for GET request
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})

}

// Availability is the search-availability page handler for GET request
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})

}

// MakeReservation is the make-reservation page handler for GET request
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{})
}

//Handler for POST request

// PostAvailability is the search-availability page handler for POST request
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	_, err := w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
	if err != nil {
		log.Fatal(err)
	}

}
