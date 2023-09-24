package main

import (
	"database/sql"
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/driver"
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

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer func(SQL *sql.DB) {
		err := SQL.Close()
		if err != nil {

		}
	}(db.SQL)

	//Config server
	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	//Start application
	log.Println("Starting application on port:", portNumber)
	_ = srv.ListenAndServe()

}

func run() (*driver.DB, error) {
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

	//Connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=2345 dbname=bookings user=postgres password=12032004")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	var tc map[string]*template.Template
	if app.UseCache {
		tc, err = render.CreateTemplate()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		app.TemplateCache = tc
	}

	//Link AppConfig to components
	render.NewTemplate(&app)
	Repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(Repo)
	helpers.NewHelpers(&app)

	return db, nil
}
