package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var functions = template.FuncMap{}

// app is the app config repo
var app *config.AppConfig

var pathToTemplates = "./templates"

// NewTemplate get App config from config to use in render
func NewTemplate(a *config.AppConfig) {
	app = a
}

// AddDefaultData get a TemplateData and return a default template data
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	//add notify
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	//add CSRFToken to template data
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate get a html page and render it
func RenderTemplate(w http.ResponseWriter, r *http.Request, html string, td *models.TemplateData) error {
	html = strings.Trim(html, " ")
	var tc map[string]*template.Template //declare template cache
	var err error
	if !app.UseCache {
		//create new template cache
		tc, err = CreateTemplate()
		if err != nil {
			return errors.New("cannot create template cache")
		}
	} else {
		//get template cache from app config
		tc = app.TemplateCache
	}
	//get requested template from cache
	t, ok := tc[html]
	if !ok {
		return errors.New(fmt.Sprintf("Have no %s page in cache \n", html))
	}
	//create new buffer memory for store template cache
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err = t.Execute(buf, td)
	if err != nil {
		log.Println(err)
		return err
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CreateTemplate create a template cache as a map,
// key: page name, value: template (html template)
func CreateTemplate() (map[string]*template.Template, error) {
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
