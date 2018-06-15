package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestStaticServerForAValidURL(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/index", nil)

	server := NewStaticServer("/index","../public/index.html").Router()
	server.ServeHTTP(w, r)

	expectedStatusCode := http.StatusOK
	actualStatusCode := w.Code

	if actualStatusCode != expectedStatusCode {
		t.Errorf("expected %d got %d", expectedStatusCode, actualStatusCode)
	}

	expectedString := "<!DOCTYPE html>"
	actualString := w.Body.String()

	if !strings.Contains(actualString, expectedString) {
		t.Errorf("expected %v contains %v", actualString, expectedString)
	}
}

func TestStaticServerForAnInvalidURL(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/invalid", nil)

	server := NewStaticServer("/index","../public/index.html").Router()
	server.ServeHTTP(w, r)

	expectedStatusCode := http.StatusNotFound
	actualStatusCode := w.Code

	if actualStatusCode != expectedStatusCode {
		t.Errorf("expected %d got %d", expectedStatusCode, actualStatusCode)
	}

	expectedString := "404 page not found\n"
	actualString := w.Body.String()

	if !strings.Contains(actualString, expectedString) {
		t.Errorf("expected %v contains %v", actualString, expectedString)
	}
}
