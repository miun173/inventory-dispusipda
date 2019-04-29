package main

import (
	"log"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/handler"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/repo"
)

func init() {
	repo.InitDB()
}

func main() {
	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "./static"
		log.Println("static dir not specified")
	} else {
		log.Println("serve static from: " + staticDir)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/people", handler.GetPeople).Methods("GET")
	router.HandleFunc("/person", handler.GetPerson).Methods("GET")
	router.HandleFunc("/person", handler.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", handler.DeletePerson).Methods("DELET")
	router.HandleFunc("/person/{id}", handler.UpdatePerson).Methods("PUT")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir + "/")))

	log.Printf("Starting server on http://localhost:%s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
