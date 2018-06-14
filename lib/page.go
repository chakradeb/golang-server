package lib

import (
	"net/http"
	"html/template"
	"fmt"
)

func PageHandler(filename string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := template.ParseFiles(filename)
		if err != nil {
			msg := fmt.Sprintf("page %s not found", filename)
			http.Error(w, msg, http.StatusNotFound)
			return
		}
		err = file.Execute(w, nil)
	}
}
