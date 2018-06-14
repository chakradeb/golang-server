package lib

import (
	"net/http"
)

func Router() *http.ServeMux {
	path := "/Users/debarc/go/src/github.com/debuc/golang-server/public/index.html"
	page := PageHandler(path)

	mux := http.NewServeMux()
	mux.HandleFunc("/index", page)
	return mux
}
