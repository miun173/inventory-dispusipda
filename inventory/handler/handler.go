package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// Ping check if service ready
func Ping(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "pong"}
	log.Print("ping")
	json.NewEncoder(w).Encode(msg)
}
