package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/williamligtenberg/workout-tracker/auth"
	db "github.com/williamligtenberg/workout-tracker/database"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 🔐 Step 1: Extract JWT from cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
		return
	}

	// 🔐 Step 2: Validate the token
	claims, err := auth.ValidateToken(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
		return
	}

	userID := r.URL.Path[len("/users/"):]
	if userID == "" {
		log.Printf("[ERROR] User ID is required")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	if claims.Subject != userID {
		http.Error(w, "Forbidden: you can only delete your own account", http.StatusForbidden)
		return
	}

	stmt, err := db.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Printf("[ERROR] Failed to prepare statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(userID)
	if err != nil {
		log.Printf("[ERROR] Failed to execute statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	log.Printf("[INFO] User deleted successfully")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"id":      userID,
		"success": "true",
		"message": "User deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}
