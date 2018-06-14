package main

import (
	"log"
	"net/http"
	"github.com/debuc/golang-server/lib"
	"os"
)

func main() {
	server := lib.Router()

	port := os.Getenv("PORT")
	log.Printf("Serving on port %v", port)
	err := http.ListenAndServe(":"+ port, server)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
