package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/handlers"
	"github.com/caovanhoang63/bookings/internal/helpers"
	"github.com/caovanhoang63/bookings/internal/models"
	"github.com/caovanhoang63/bookings/internal/render"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	//run

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	//Config server
	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	//Start application
	log.Println("Starting application on port:", portNumber)
	_ = srv.ListenAndServe()

}

func run() error {
	//what put in the session`
	gob.Register(models.Reservation{})

	//config application
	app.UseCache = true

	//config logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	//change this to true when in production
	app.InProduction = false

	//config session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = app.InProduction
	session.Cookie.Persist = true //the cookie stills alive when close browser tab
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		tc, err = render.CreateTemplate()
		if err != nil {
			return err
		}
		app.TemplateCache = tc
	}

	//Link AppConfig to components
	render.NewTemplate(&app)
	Repo := handlers.NewRepo(&app)
	handlers.NewHandlers(Repo)
	helpers.NewHelpers(&app)

	return nil
}
