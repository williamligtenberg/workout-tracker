package handlers

import "net/http"

func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("user_id")
	w.Write([]byte("User ID: " + userID))
}
