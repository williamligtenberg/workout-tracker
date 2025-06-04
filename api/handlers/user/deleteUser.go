package user

import (
	"encoding/json"
	"log"
	"net/http"
	db "workout-tracker/api/database"
	models "workout-tracker/api/models"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Path[len("/users/"):]

	if userID == "" {
		log.Printf("[ERROR] User ID is required")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var user models.User

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

	log.Printf("[INFO] User deleted successfully: %d", user.Id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "User deleted successfully",
		"user":    user.Username,
	}
	json.NewEncoder(w).Encode(response)
}
