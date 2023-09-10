package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/forms"
	"github.com/caovanhoang63/bookings/internal/models"
	"github.com/caovanhoang63/bookings/internal/render"
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
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation handles the posing of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

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

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		log.Println(err.Error())
	}

}
