package main

import (
	"github.com/caovanhoang63/bookings/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	var a config.AppConfig

	h := routes(&a)

	switch v := h.(type) {
	case *chi.Mux:
		//do nothing, test passed
	default:
		t.Errorf("type not is *chi.Mux, type is %T", v)
	}
}
