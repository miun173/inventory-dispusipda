package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/repo"
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

func GetAllBarangKeluar(w http.ResponseWriter, r *http.Request) {
	brgs, err := repo.GetAllBarangKeluar()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(brgs)
}

// GetJurnal list all available jurnal
func GetJurnal(w http.ResponseWriter, r *http.Request) {
	var jurnals []models.Jurnal
	brgs, err := repo.GetAllBarang()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	for _, b := range brgs {
		penyusutan := (b.Harga - b.NilaiSisa) / float64(b.UmurEkonomis)
		j := models.Jurnal{
			Barang:     b,
			Penyusutan: penyusutan,
		}
		jurnals = append(jurnals, j)
	}

	jData, err := json.Marshal(jurnals)
	if err != nil {
		log.Printf("error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "something bad"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jData)
}
