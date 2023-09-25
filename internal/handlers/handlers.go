package handlers

import (
	"encoding/json"
	"errors"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/driver"
	"github.com/caovanhoang63/bookings/internal/forms"
	"github.com/caovanhoang63/bookings/internal/helpers"
	"github.com/caovanhoang63/bookings/internal/models"
	"github.com/caovanhoang63/bookings/internal/render"
	"github.com/caovanhoang63/bookings/internal/repository"
	"github.com/caovanhoang63/bookings/internal/repository/dbrepo"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// convert date's layout to dd/mm/yyyy
const dateTimeLayout = "02-01-2006"

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
	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("cannot get data from session"))
		return
	}

	sd := res.StartDate.Format(dateTimeLayout)
	ed := res.EndDate.Format(dateTimeLayout)

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// PostReservation handles the posing of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	startDate, err := time.Parse(dateTimeLayout, sd)
	if err != nil {
		helpers.ServerError(w, err)
	}

	endDate, err := time.Parse(dateTimeLayout, ed)
	if err != nil {
		helpers.ServerError(w, err)
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restriction := models.RoomRestrictions{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        roomID,
		RestrictionID: 1,
		ReservationID: newReservationID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "reservation-summary", http.StatusSeeOther)

}

// ReservationSummary is the reservation-summary page handler for render the reservation summary
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

	startDate, err := time.Parse(dateTimeLayout, start)
	if err != nil {
		helpers.ServerError(w, err)
	}
	endDate, err := time.Parse(dateTimeLayout, end)
	if err != nil {
		helpers.ServerError(w, err)
	}
	rooms, err := m.DB.SearchAvailabilityForAllRoom(startDate, endDate)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(rooms) == 0 {
		m.App.Session.Put(r.Context(), "error", "No availability")
		http.Redirect(w, r, "search-availability", http.StatusSeeOther)
		return
	}

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	m.App.Session.Put(r.Context(), "reservation", res)
	m.App.Session.Put(r.Context(), "rooms", rooms)

	render.Template(w, r, "choose-room.page.html", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, ok := m.App.Session.Get(r.Context(), "rooms").([]models.Room)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	var room models.Room
	for _, room = range rooms {
		if room.ID == roomID {
			break
		}
	}

	res, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID
	res.Room = room

	m.App.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/make-reservation", http.StatusSeeOther)

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
