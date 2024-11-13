package user

import (
	"encoding/json"
	"log"
	"net/http"
	models "workout-tracker/api/models"
	api "workout-tracker/api/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIDString := r.PathValue("user_id")
	userID, err := api.StringToInt(userIDString)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	user := models.User{
		Id:       userID,
		Username: "",
	}

	w.WriteHeader(http.StatusOK)

	log.Printf("[INFO] Retrieved user | user_id:%d | method:%s | path:%s", userID, r.Method, r.URL.Path)

	if err := json.NewEncoder(w).Encode(struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}{
		Id:   user.Id,
		Name: user.Username,
	}); err != nil {
		log.Printf("[ERROR] Failed to encode user response | user_id=%d | method=%s | error=%v", userID, r.Method, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
