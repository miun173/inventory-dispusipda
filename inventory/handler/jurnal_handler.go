package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/miun173/inventory-dispusibda/inventory/app"
)

// JurnalHandler handler for jurnal routes
type JurnalHandler interface {
	GetJurnal(w http.ResponseWriter, r *http.Request)
}

type jurnalHandler struct {
	jurnalService app.JurnalService
}

func NewJurnalHandler(jurnalService app.JurnalService) JurnalHandler {
	return &jurnalHandler{jurnalService}
}

// GetJurnal list all available jurnal/inventaris
func (j *jurnalHandler) GetJurnal(w http.ResponseWriter, r *http.Request) {
	// inisiasi
	jurnals, err := j.jurnalService.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "system error"})
		return
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
