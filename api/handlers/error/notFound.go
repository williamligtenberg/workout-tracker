package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	models "workout-tracker/api/models"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	log.Printf("[ERROR] 404 - Not Found | Method: %s | Path: %s", r.Method, r.URL.Path)

	response := models.ErrorResponse{
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Message: "The requested resource could not be found.",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("[ERROR] Failed to encode 404 response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
