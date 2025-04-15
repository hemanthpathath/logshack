package router

import (
	"net/http"

	"github.com/hemanthpathath/logsnag/logsnag-backend/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Define your routes here
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/events", handler.EventsHandler)
	return mux
}
