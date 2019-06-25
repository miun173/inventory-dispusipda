package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/repository"
)

type BarangHandler interface {
	Save(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type barangHandler struct {
	brgRepo repository.BarangRepo
}

// NewBarangHandler get new barang handler object
func NewBarangHandler(brgRepo repository.BarangRepo) BarangHandler {
	return &barangHandler{brgRepo}
}

// Save handle create barang
func (b *barangHandler) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// barang diambil dari inputan
	var brg models.Barang
	_ = json.NewDecoder(r.Body).Decode(&brg)

	err := b.brgRepo.Save(&brg)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	json.NewEncoder(w).Encode(brg)
}

// Get handle get all barang
func (b *barangHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
		return
	}

	brg, err := b.brgRepo.Get(id)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	json.NewEncoder(w).Encode(brg)
}

// Update handle update for a barang
func (b *barangHandler) Update(w http.ResponseWriter, r *http.Request) {
	var brg models.Barang
	err := json.NewDecoder(r.Body).Decode(&brg)
	if err != nil {
		log.Printf("%+v", err)
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	err = b.brgRepo.Update(brg)

	json.NewEncoder(w).Encode(brg)
}

// GetAll list all available barang keluar
func (b *barangHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	brgs, err := b.brgRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(brgs)
}
