package lib

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticServer_Router(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/index", nil)

	server := NewStaticServer("/index","../public/index.html").Router()
	server.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code, "Invalid response code")
	assert.Contains(t, w.Body.String(), "<!DOCTYPE html>", "Invalid response body")
}

func TestStaticServer_Router2(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/invalid", nil)

	server := NewStaticServer("/index","../public/index.html").Router()
	server.ServeHTTP(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code, "Invalid response code")
	assert.NotContains(t, w.Body.String(), "<!DOCTYPE html>", "Invalid response body")
	assert.Equal(t, w.Body.String(), "404 page not found\n", "Invalid Error Message")
}
