package render

import (
	"bytes"
	"fmt"
	"github.com/caovanhoang63/bookings/pkg/config"
	"github.com/caovanhoang63/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// app is the app config repo
var app *config.AppConfig

// NewTemplate get App config from config to use in render
func NewTemplate(a *config.AppConfig) {
	app = a
}

// AddDefaultData get a TemplateData and return a default template data
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	html = strings.Trim(html, " ")
	var tc map[string]*template.Template //declare template cache
	var err error
	if !app.UseCache {
		//create new template cache
		tc, err = CreateTemplate()
		if err != nil {
			log.Fatal("cannot create template cache")
		}
	} else {
		//get template cache from app config
		tc = app.TemplateCache
	}
	//get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatalf("Have no %s page in cache", html)
	}
	//create new buffer memory for store template cache
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err = t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateTemplate() (map[string]*template.Template, error) {
	log.Print("Create template cache conclude: ")
	templateCache := make(map[string]*template.Template)
	//get all the file named *page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	fmt.Println(pages)
	if err != nil {

		return templateCache, err
	}
	//loop for all html page in pages
	for _, page := range pages {
		//get name of html page
		name := filepath.Base(page)
		//Parse html file
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}
		//Load all dependency layout
		var matches []string
		matches, err = filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			//Parse all template layout that the page needs
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return templateCache, err
			}
		}
		//append template to cache
		fmt.Println(name)
		templateCache[name] = ts
	}
	return templateCache, nil
}
