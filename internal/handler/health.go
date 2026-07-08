package handler

import (
	"encoding/json"
	"net/http"
)

// HealthHandler exposes a simple health-check endpoint.
type HealthHandler struct{}

// Check returns the server health status.
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
