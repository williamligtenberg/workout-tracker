package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/williamligtenberg/workout-tracker/auth"
	models "github.com/williamligtenberg/workout-tracker/models"
	"github.com/williamligtenberg/workout-tracker/utils"

	"log"

	db "github.com/williamligtenberg/workout-tracker/database"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("[ERROR] Decoding body: %v", err)
		utils.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		log.Printf("[ERROR] Missing required fields in user data")
		utils.JSONError(w, http.StatusBadRequest, "Username, email, and password are required")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ERROR] Failed to hash password: %v", err)
		utils.JSONError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	user.Password = string(hashedPassword)

	stmt, err := db.DB.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("[ERROR] Failed to prepare statement: %v", err)
		utils.JSONError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, hashedPassword)
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
