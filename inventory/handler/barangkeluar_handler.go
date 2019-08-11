package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/miun173/inventory-dispusibda/inventory/models"
	"github.com/miun173/inventory-dispusibda/inventory/repository"
)

// BarangKeluarHandler interface
type BarangKeluarHandler interface {
	Save(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type barangKeluarHandler struct {
	brgRepo       repository.BarangRepo
	brgKeluarRepo repository.BarangKeluarRepo
}

// NewBarangKeluarHandler get new barang handler object
func NewBarangKeluarHandler(
	brgRepo repository.BarangRepo,
	brgKeluarRepo repository.BarangKeluarRepo,
) BarangKeluarHandler {
	return &barangKeluarHandler{brgRepo, brgKeluarRepo}
}

// Save handle barang keluar
func (b *barangKeluarHandler) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var brg models.BarangKeluar
	_ = json.NewDecoder(r.Body).Decode(&brg)

	isBarangExists, err := b.brgRepo.IsExists(brg.BarangID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	if !isBarangExists {
		errMsg := fmt.Sprintf("barang with barangID %d is not exists", brg.BarangID)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": errMsg})
		return
	}

	// create barang if barangID exist
	err = b.brgKeluarRepo.Save(&brg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(brg)
}

// GetAll list all available barang keluar
func (b *barangKeluarHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	brgs, err := b.brgKeluarRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(brgs)
}
