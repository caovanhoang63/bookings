package render

import (
	"github.com/caovanhoang63/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"

	tc, err := CreateTemplate()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = tc
	var r *http.Request

	r, err = getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myResponseWriter

	err = Template(&ww, r, "home.page.html", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}
	err = Template(&ww, r, "non-exist.page.html", &models.TemplateData{})
	if err == nil {
		t.Error("rendered a template that not exist")
	}
}

func TestNewTemplate(t *testing.T) {
	NewRenderer(app)
}

func TestCreateTemplate(t *testing.T) {
	_, err := CreateTemplate()
	if err != nil {
		t.Error(err)
	}
}

// getSession return a http Request with a session data
// getSession is responsible as SessionLoad middleware
func getSession() (*http.Request, error) {

	r, err := http.NewRequest("GET", "/test-url", nil)
	if err != nil {
		return r, err
	}

	ctx := r.Context()

	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))

	r = r.WithContext(ctx)

	return r, nil
}
