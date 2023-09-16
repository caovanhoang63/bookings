package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)

	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid ")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)

	form := New(r.PostForm)
	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required missing")
	}

	dataPost := url.Values{}
	dataPost.Add("a", "a")
	dataPost.Add("b", "b")
	dataPost.Add("c", "c")

	r = httptest.NewRequest("POST", "/any", nil)
	r.PostForm = dataPost
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("form shows invalid when all having all required field")
	}
}

//func TestNew(t *testing.T) {
//	r := httptest.NewRequest("POST", "/any", nil)
//
//	form := New(r.PostForm)
//
//
//}

func TestHas(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)

	if Has("test_field", r) {
		t.Error("the request has no test_field but the function that shows it has")
	}

	dataPost := url.Values{}
	dataPost.Add("test_field", "test")
	r = httptest.NewRequest("POST", "/any", nil)
	r.Form = dataPost

	if !Has("test_field", r) {
		t.Error("the request has test_field but the function shows that it doesn't have")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	valid := "cao@gmail.com"
	dataPost := url.Values{}
	dataPost.Add("Email", valid)
	r.PostForm = dataPost
	form := New(r.PostForm)

	if !form.IsEmail("Email") {
		t.Error("form shows invalid when values valid")
	}

	invalid := "cao"
	dataPost.Set("Email", invalid)
	r.PostForm = dataPost
	form = New(r.PostForm)
	if form.IsEmail("Email") {
		t.Error("form shows valid when values invalid")
	}
}

func TestForm_IsPhoneNumber(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	valid := "0896374872"
	dataPost := url.Values{}
	dataPost.Add("Phone", valid)
	r.PostForm = dataPost
	form := New(r.PostForm)

	if !form.IsPhoneNumber("Phone") {
		t.Error("form shows invalid when values valid")
	}

	invalid := "11"
	dataPost.Set("Phone", invalid)
	r.PostForm = dataPost
	form = New(r.PostForm)
	if form.IsPhoneNumber("Phone") {
		t.Error("form shows valid when values invalid")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/any", nil)
	valid := "333"
	dataPost := url.Values{}
	dataPost.Add("test_field", valid)
	r.PostForm = dataPost
	form := New(r.PostForm)

	if !form.MinLength("test_field", 3) {
		t.Error("form shows invalid when values valid")
	}

	invalid := "22"
	dataPost.Set("test_field", invalid)
	r.PostForm = dataPost
	form = New(r.PostForm)
	if form.MinLength("test_field", 3) {
		t.Error("form shows valid when values invalid")
	}
}
