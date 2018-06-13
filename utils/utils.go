package utils

import "net/http"

type Handler interface {
	HandleFunction(w http.ResponseWriter,r http.Request)
}