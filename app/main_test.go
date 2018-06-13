package main

import (
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

type MockPage struct {
	mock.Mock
}

func (page MockPage) RenderHTML(res http.ResponseWriter,filename string) {
	res.Write([]byte(filename))
}

func TestLandingPageHandler(t *testing.T) {
	mock := MockPage{}
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	Render(mock,"abc")(res,req)
	assert.Equal(t,"abc",res.Body.String())
}
