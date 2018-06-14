package lib

import (
	"net/http"
)

type Route interface {
	Router() *http.ServeMux
}

type StaticServer struct {
	route string
	filepath string
}

func NewStaticServer(route string, filepath string) *StaticServer {
	return &StaticServer{route, filepath}
}

func (p *StaticServer) Router() *http.ServeMux {
	page := NewPage(p.filepath)

	mux := http.NewServeMux()
	mux.HandleFunc(p.route, page.Handle())
	return mux
}
