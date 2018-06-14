package lib

import (
	"net/http"
	"html/template"
	"fmt"
)

type Handler interface {
	Handle() func(w http.ResponseWriter, r *http.Request)
}

type Page struct {
	path string
}

func NewPage(path string) (p *Page) {
	return &Page{path}
}

func (p *Page) Handle() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := template.ParseFiles(p.path)
		if err != nil {
			msg := fmt.Sprintf("page %s not found", p.path)
			http.Error(w, msg, http.StatusNotFound)
			return
		}
		err = file.Execute(w, nil)
	}
}
