package router

import (
	"net/http"

	"github.com/hemanthpathath/logshack/logshack-backend/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Define your routes here
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/events", handler.CreateEvent)

	return mux
}
