package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/miun173/inventory-dispusibda/cmd/inventory/repository"

	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
)

// RkbmdHandler interface
type RkbmdHandler interface {
	Save(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type rkbmdHandler struct {
	rkbmdRepo repository.RkbmdRepo
}

// NewRkbmdHandler get new barang handler object
func NewRkbmdHandler(rkbdmRepo repository.RkbmdRepo) RkbmdHandler {
	return &rkbmdHandler{rkbdmRepo}
}

// Save handler
func (rh *rkbmdHandler) Save(w http.ResponseWriter, r *http.Request) {
	// create Rkbmd
	now := time.Now().Unix() * 1000
	rkbmd := models.Rkbmd{
		TglBuat: int(now),
	}
	err := rh.rkbmdRepo.Save(&rkbmd)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "something bad"})
	}

	// Create Rkbmd detail
	var dRkbmd []models.DetailRkbmd
	mapData := map[string][]models.DetailRkbmd{"rkbmd": dRkbmd}
	err = json.NewDecoder(r.Body).Decode(&mapData)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "something bad"})
		return
	}

	err = rh.rkbmdRepo.SaveDetail(mapData["rkbmd"], rkbmd.ID)
	if err != nil {
		log.Printf("%+v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "something bad"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"rkbmdID": rkbmd.ID})
}

// GetAll handler
func (rh *rkbmdHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rkbmds, err := rh.rkbmdRepo.GetAll()
	if err != nil {
		log.Printf("%+v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "something bad"})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rkbmds)
}

// Update handler
func (rh *rkbmdHandler) Update(w http.ResponseWriter, r *http.Request) {
	var rkbmdDetail models.RkbmdDetail
	err := json.NewDecoder(r.Body).Decode(&rkbmdDetail)
	if err != nil {
		log.Printf("%+v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "something bad"})
		return
	}

	// update rkbmd detail
	err = rh.rkbmdRepo.UpdateDetail(rkbmdDetail)
	if err != nil {
		log.Printf("%+v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "something bad"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "updated"})
}
