package handlers

import "net/http"

func UserHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("User ID: " + userID))
}
