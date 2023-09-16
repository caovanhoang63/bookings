package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/caovanhoang63/bookings/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})

	//config application
	testApp.UseCache = true

	//change this to true when in production
	testApp.InProduction = false

	//config session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = testApp.InProduction
	session.Cookie.Persist = true //the cookie stills alive when close browser tab
	session.Cookie.SameSite = http.SameSiteLaxMode

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myResponseWriter struct{}

//Header() Header
//Write([]byte) (int, error)
//WriteHeader(statusCode int)

func (mr *myResponseWriter) Header() http.Header {
	var h http.Header
	return h
}

func (mr *myResponseWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func (mr *myResponseWriter) WriteHeader(_ int) {
}
