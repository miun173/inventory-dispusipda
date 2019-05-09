package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"

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

	router.HandleFunc("/api/ping", handler.Ping).Methods("GET")

	// users
	router.HandleFunc("/api/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/users", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users", handler.CreateUser).Methods("POST")

	// barang
	router.HandleFunc("/api/barang/{id}", handler.GetBarang).Methods("GET")
	router.HandleFunc("/api/barang", handler.GetAllBarang).Methods("GET")
	router.HandleFunc("/api/barang", handler.CreateBarang).Methods("POST")
	router.HandleFunc("/api/barang-keluar", handler.CreateBarangKeluar).Methods("POST")
	router.HandleFunc("/api/barang-keluar", handler.GetAllBarangKeluar).Methods("GET")

	// jurnal
	router.HandleFunc("/api/jurnal", handler.GetJurnal).Methods("GET")

	// frontend
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir + "/")))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	log.Printf("Starting server on http://localhost:%s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, n))
}
