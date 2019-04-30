package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/repo"
	"log"
	"net/http"
	"strconv"
)

// var people []models.Person

// Ping check if service ready
func Ping(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "pong"}
	log.Print("ping")
	json.NewEncoder(w).Encode(msg)
}

// CreateUser handle user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := repo.CreateUser(&user)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetAllUsers handle get all users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repo.GetAllUser()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Login handle user login
func Login(w http.ResponseWriter, r *http.Request) {

}

// CreateBarang handle create barang
func CreateBarang(w http.ResponseWriter, r *http.Request) {
	var brg models.Barang
	_ = json.NewDecoder(r.Body).Decode(&brg)
	err := repo.CreateBarang(&brg)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(brg)
}

// GetBarang handle get all barang
func GetBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
		return
	}

	var brg models.Barang
	err = repo.GetBarang(id, &brg)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(brg)
}

// CreateBarangKeluar handle barang keluar
func CreateBarangKeluar(w http.ResponseWriter, r *http.Request) {
	var brg models.BarangKeluar
	_ = json.NewDecoder(r.Body).Decode(&brg)
	err := repo.CreateBarangKeluar(&brg)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(brg)
}

// GetJurnal get jurnal
func GetJurnal(w http.ResponseWriter, r *http.Request) {
	brgs, err := repo.GetAllBarang()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(brgs)
}
