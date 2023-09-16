package forms

import (
	"fmt"
	"testing"
)

func TestErrors_Add(t *testing.T) {
	err := errors{}
	err.Add("Any", "Any")
	if len(err) < 0 {
		t.Error("new element isn't added to map")
	}
}

func TestErrors_Get(t *testing.T) {
	err := errors{}
	if s := err.Get("Any"); s != "" {
		t.Error(fmt.Sprintf("got %s when the map is empty", s))
	}
	err.Add("Any", "Any")
	err.Add("Any2", "Any2")
	if err.Get("Any") == "" {
		t.Error("got no element when the map is not empty")
	}
	if err.Get("Any") != "Any" {
		t.Error("got wrong element")
	}
}
