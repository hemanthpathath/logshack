package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/hemanthpathath/logshack/logshack-backend/internal/db"
)

type Event struct {
	UserID    string                 `json:"user_id"`
	ProjectID string                 `json:"project_id"`
	EventType string                 `json:"event_type"`
	Payload   map[string]interface{} `json:"payload"`
}

func (e *Event) Validate() error {
	if e.UserID == "" || e.ProjectID == "" || e.EventType == "" {
		return errors.New("missing required fields")
	}
	return nil
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var e Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	if err := e.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Pool.Exec(r.Context(),
		"INSERT INTO events (user_id, project_id, event_type, payload) VALUES ($1, $2, $3, $4)", e.UserID, e.ProjectID, e.EventType, e.Payload)

	if err != nil {
		log.Printf("insert event failed: %v", err)
		http.Error(w, "Failed to insert event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Event created"})
}
