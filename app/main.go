package main

import (
	"net/http"
	"html/template"
	"github.com/tkanos/gonfig"
	"fmt"
	"strconv"
	"os"
)

type Renderer interface {
	RenderHTML(w http.ResponseWriter,filename string)
}

type Configuration struct {
	Port int
	Filename string
}

var config = Configuration{}

type Page struct {
	content []byte
}

func (page Page)RenderHTML(res http.ResponseWriter,filename string) {
	file, err := template.ParseFiles( filename + ".html")
	if err!=nil {
		page.content = []byte("page not found")
		res.WriteHeader(404)
		res.Write(page.content)
	}
	file.Execute(res,page.content)
}

func Render(page Renderer,filename string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter,r *http.Request) {
		page.RenderHTML(w,filename)
	}

}

func LogPort(Port int)  {
	fmt.Println("Serving on Port "+ strconv.Itoa(Port))
}

func main() {
	gonfig.GetConf(os.Getenv("CONFIG_PATH"),&config)
	http.HandleFunc("/", Render(Page{},config.Filename))
	LogPort(config.Port)
	http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
}
