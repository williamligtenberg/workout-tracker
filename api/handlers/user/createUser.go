package user

import (
	"encoding/json"
	"net/http"
	"time"
	"workout-tracker/api/auth"
	models "workout-tracker/api/models"

	"log"
	db "workout-tracker/api/database"

	"github.com/go-sql-driver/mysql"
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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			log.Printf("[WARN] Duplicate username or email: %v", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{
				"success": "false",
				"error":   "Username or email already exists",
			})
			return
		}
		log.Printf("[ERROR] Failed to execute statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tokenString, err := auth.GenerateToken(user.Username)
	if err != nil {
		log.Printf("[ERROR] Failed to generate JWT token: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	log.Printf("[INFO] User created successfully: %s", user.Username)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"success": "true",
		"payload": "User created successfully",
		"id":      user.Username,
	}
	json.NewEncoder(w).Encode(response)
}
