package main

import (
	"log"
	"net/http"

	"github.com/hemanthpathath/logsnag/logsnag-backend/internal/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
