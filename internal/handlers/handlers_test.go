package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"ge", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"sa-post", "/search-availability", "POST", []postData{
		{"start", "10-10-2022"},
		{"end", "20-10-2022"},
	}, http.StatusOK},
	{"sa-post-json", "/search-availability", "POST", []postData{
		{"start", "10-10-2022"},
		{"end", "20-10-2022"},
	}, http.StatusOK},
	{"mk-reservation", "/make-reservation", "POST", []postData{
		{"first_name", "Cao"},
		{"last_name", "Hoang"},
		{"email", "caovanhoang204@gmail.com"},
		{"phone", "0896374872"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	testRoutes := getRoutes()
	//create test server
	ts := httptest.NewTLSServer(testRoutes)
	defer ts.Close()
	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d ", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		} else {
			values := url.Values{}
			for _, v := range e.params {
				values.Add(v.key, v.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)

			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d ", e.name, e.expectedStatusCode, resp.StatusCode)
			}

		}
	}
}
