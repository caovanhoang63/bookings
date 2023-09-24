package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/driver"
	"github.com/caovanhoang63/bookings/internal/forms"
	"github.com/caovanhoang63/bookings/internal/helpers"
	"github.com/caovanhoang63/bookings/internal/models"
	"github.com/caovanhoang63/bookings/internal/render"
	"github.com/caovanhoang63/bookings/internal/repository"
	"github.com/caovanhoang63/bookings/internal/repository/dbrepo"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
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
	render.Template(w, r, "home.page.html", &models.TemplateData{})
}

// About is the about page handler for GET request
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.html", &models.TemplateData{})
}

// Contact is the contact page handler for GET request
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "contact.page.html", &models.TemplateData{})

}

// Generals is the generals page handler for GET request
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "generals.page.html", &models.TemplateData{})

}

// Major is the major page handler for GET request
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "majors.page.html", &models.TemplateData{})

}

// Availability is the search-availability page handler for GET request
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "search-availability.page.html", &models.TemplateData{})

}

// MakeReservation is the make-reservation page handler for GET request
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posing of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.IsEmail("email")
	form.IsPhoneNumber("phone")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "reservation-summary", http.StatusSeeOther)

}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation form session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
}

//Handler for POST request

// PostAvailability is the search-availability page handler for POST request
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	_, err := w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
	if err != nil {
		helpers.ServerError(w, err)
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
		helpers.ServerError(w, err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		helpers.ServerError(w, err)
	}

}
