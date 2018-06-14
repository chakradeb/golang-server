package mocks

import (
	"github.com/stretchr/testify/mock"
)

type Page struct {
	mock.Mock
}

func (m *Page) Handle() interface{} {
	args := m.Called()
	return args.Get(0)
}
