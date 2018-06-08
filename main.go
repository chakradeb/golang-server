package main

import (
	"net/http"
	"html/template"
	"github.com/tkanos/gonfig"
	"fmt"
	"strconv"
	"os"
)

type Configuration struct {
	Port int
	Filename string
}

type Page struct {
	content []byte
}

var config = Configuration{}

func RenderHTML(res http.ResponseWriter,filename string,p *Page) {
	file, _ := template.ParseFiles(filename + ".html")
	file.Execute(res,p.content)
}

func LandingPageHandler(res http.ResponseWriter, req *http.Request)  {
	landingPage := &Page{}
	RenderHTML(res,config.Filename,landingPage)
}

func LogPort(Port int)  {
	fmt.Println("Serving on Port "+ strconv.Itoa(Port))
}

func main() {
	gonfig.GetConf(os.Getenv("CONFIG_PATH"),&config)
	fmt.Println(config.Port)
	fmt.Println(config.Filename)
	http.HandleFunc("/", LandingPageHandler)
	LogPort(config.Port)
	http.ListenAndServe(":" + strconv.Itoa(config.Port), nil)
}
