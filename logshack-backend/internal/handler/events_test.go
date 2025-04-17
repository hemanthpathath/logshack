package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hemanthpathath/logshack/logshack-backend/internal/db"
)

func TestMain(m *testing.M) {
	db.InitPool("postgres://postgres:postgres@localhost:5432/logshack?sslmode=disable")
	code := m.Run()
	os.Exit(code)
}

func TestCreateEvent_Valid(t *testing.T) {
	payload := []byte(`{
		"user_id": "TUser1",
		"project_id": "TProject1",
		"event_type": "TSignup",
		"payload": { "plan": "pro" }
	}`)

	req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	CreateEvent(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rr.Code)
	}
}

func TestCreateEvent_InvalidJson(t *testing.T) {
	payload := []byte(`{ "user_id": "TUser1" }`)

	req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	CreateEvent(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
