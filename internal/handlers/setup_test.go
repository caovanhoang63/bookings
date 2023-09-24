package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/driver"
	"github.com/caovanhoang63/bookings/internal/models"
	"github.com/caovanhoang63/bookings/internal/render"
	"github.com/go-chi/chi/v5"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

//this file has all set up environment for handlers testing

var app config.AppConfig
var session *scs.SessionManager
var functions template.FuncMap
var pathToTemplates = "./../../templates"
var infoLog *log.Logger
var errorLog *log.Logger

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the sessions on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func getRoutes() http.Handler {

	//what put in the session`
	gob.Register(models.Reservation{})

	//config application
	app.UseCache = true

	//change this to true when in production
	app.InProduction = false

	//config session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = app.InProduction
	session.Cookie.Persist = true //the cookie stills alive when close browser tab
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session
	//Connect to database
	db, err := driver.ConnectSQL("host=localhost port=2345 dbname=bookings user=postgres password=12032004")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	//config logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
	var tc map[string]*template.Template
	if app.UseCache {
		tc, err = CreateTestTemplate()
		if err != nil {
			log.Fatal(err)
		}
		app.TemplateCache = tc
	}

	//Link AppConfig to components
	render.NewTemplate(&app)
	Repo := NewRepo(&app, db)
	NewHandlers(Repo)

	mux := chi.NewMux()

	//use middleware
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//use Handlers

	mux.Get("/", Repo.Home)

	mux.Get("/about", Repo.About)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/generals-quarters", Repo.Generals)

	mux.Get("/majors-suite", Repo.Major)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Get("/search-availability-json", Repo.AvailabilityJSON)

	mux.Get("/make-reservation", Repo.MakeReservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	//File server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

func CreateTestTemplate() (map[string]*template.Template, error) {
	templateCache := make(map[string]*template.Template)
	//get all the file named *page.html from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {

		return templateCache, err
	}
	//loop for all html page in pages
	for _, page := range pages {
		//get name of html page
		name := filepath.Base(page)
		//Parse html file
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}
		//Load all dependency layout
		var matches []string
		matches, err = filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			//Parse all template layout that the page needs
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return templateCache, err
			}
		}
		//append template to cache
		templateCache[name] = ts
	}
	return templateCache, nil
}
