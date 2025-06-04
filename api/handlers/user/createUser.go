package user

import (
	"encoding/json"
	"net/http"
	models "workout-tracker/api/models"

	"log"
	db "workout-tracker/api/database"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("[ERROR] Failed to decode request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		log.Printf("[ERROR] Missing required fields in user data")
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	stmt, err := db.DB.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("[ERROR] Failed to prepare statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("[ERROR] Failed to execute statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] User created successfully: %s", user.Username)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "User created successfully",
		"user":    user.Username,
	}
	json.NewEncoder(w).Encode(response)
}
