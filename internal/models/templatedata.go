package models

import "github.com/caovanhoang63/bookings/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string //a flash notice
	Warning   string
	Error     string
	Form      *forms.Form
}
