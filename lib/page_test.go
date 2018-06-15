package lib

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"strings"
)

func TestPage_HandleForValidFilepath(t *testing.T) {
	page := &Page{"../public/index.html"}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	page.Handle()(w,r)

	expectedString := "<!DOCTYPE html>"
	actualString := w.Body.String()

	if !strings.Contains(actualString, expectedString) {
		t.Errorf("expected %v contains %v", actualString, expectedString)
	}
}

func TestPage_HandleForInvalidFilepath(t *testing.T) {
	page := &Page{"../public/invalid.html"}

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	page.Handle()(w,r)

	expectedStatusCode := http.StatusNotFound
	actualStatusCode := w.Code

	if actualStatusCode != expectedStatusCode {
		t.Errorf("expected %d got %d", expectedStatusCode, actualStatusCode)
	}

	expectedString := "page ../public/invalid.html not found\n"
	actualString := w.Body.String()

	if !strings.Contains(actualString, expectedString) {
		t.Errorf("expected %v contains %v", actualString, expectedString)
	}
}