package handler

import (
	"encoding/json"
	"net/http"

	"github.com/miun173/inventory-dispusibda/inventory/models"
	"github.com/miun173/inventory-dispusibda/inventory/repo"
)

// CreateUser handle user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := repo.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetAllUsers handle get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repo.GetAllUser()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	json.NewEncoder(w).Encode(users)
}
