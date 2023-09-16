package main

import (
	"net/http"
	"os"
	"testing"
)

// TestMain runs before all other tests
// and in this function, all setup for testing are prepared
func TestMain(m *testing.M) {

	//Program exits with code is returned by m.Run()
	os.Exit(m.Run())
}

// my Handler implements http.Handler which ís an interface
type myHandler struct{}

// ServerHTTP implements ServeHTTP from http.Handler
// so myHandler also considered as a http.Handler
func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
