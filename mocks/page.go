package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type Page struct {
	mock.Mock
}

func (m *Page) RenderHTML(w http.ResponseWriter, filename string) {
	m.Called(w, filename)
}
