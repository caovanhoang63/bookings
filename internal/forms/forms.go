package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Error errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form field is in post and not empty
func Has(field string, r *http.Request) bool {
	return r.Form.Get(field) != ""
}
