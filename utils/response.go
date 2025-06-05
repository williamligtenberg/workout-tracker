package utils

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"success": "false",
		"error":   msg,
	})
}

func JSONSuccess(w http.ResponseWriter, status int, payload map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	payload["success"] = "true"
	json.NewEncoder(w).Encode(payload)
}
