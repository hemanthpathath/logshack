package main

import (
	"log"
	"net/http"

	"github.com/hemanthpathath/logshack/logshack-backend/internal/db"
	"github.com/hemanthpathath/logshack/logshack-backend/internal/router"
)

func main() {
	db.Init()
	r := router.NewRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
