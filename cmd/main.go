package main

import (
	"log"
	"net/http"
	"os"

	"github.com/debuc/golang-server/lib"
)

func main() {
	server := lib.NewStaticServer("/index","public/index.html")

	port := os.Getenv("PORT")
	log.Printf("Serving on port %v", port)
	err := http.ListenAndServe(":"+ port, server.Router())
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
