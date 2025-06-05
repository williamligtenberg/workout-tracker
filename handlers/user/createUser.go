package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/williamligtenberg/workout-tracker/auth"
	models "github.com/williamligtenberg/workout-tracker/models"
	"github.com/williamligtenberg/workout-tracker/utils"

	"log"

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
	user.UUID = uuid.NewString()

	err = models.CreateUser(&user)
	if err != nil {
		if models.IsDuplicateUserError(err) {
			utils.JSONError(w, http.StatusConflict, "Username or email already exists")
			return
		}
		log.Printf("[ERROR] Creating user: %v", err)
		utils.JSONError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	token, err := auth.GenerateToken(user.UUID)
	if err != nil {
		log.Printf("[ERROR] Failed to generate JWT token: %v", err)
		utils.JSONError(w, http.StatusInternalServerError, "Internal error")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	utils.JSONSuccess(w, http.StatusCreated, map[string]string{
		"message": "User created successfully",
		"id":      user.Username,
	})
}
