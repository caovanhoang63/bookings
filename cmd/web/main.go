package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/handlers"
	"github.com/caovanhoang63/bookings/internal/render"
	"html/template"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
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

	var tc map[string]*template.Template
	var err error
	if app.UseCache {

		tc, err = render.CreateTemplate()
		if err != nil {
			log.Fatal(err)
		}
		app.TemplateCache = tc
	}

	//Link AppConfig to components
	render.NewTemplate(&app)
	Repo := handlers.NewRepo(&app)
	handlers.NewHandlers(Repo)

	//Config server
	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	//Start application
	log.Println("Starting application on port:", portNumber)
	_ = srv.ListenAndServe()

}
