package main

import (
	"log"
	"net/http"

	"github.com/miun173/inventory-dispusibda/inventory/app"
	"github.com/miun173/inventory-dispusibda/inventory/handler"
	"github.com/miun173/inventory-dispusibda/inventory/infra"
	"github.com/miun173/inventory-dispusibda/inventory/repository"
	"github.com/urfave/negroni"

	"os"

	"github.com/gorilla/mux"
)

func main() {
	db := infra.NewDB()
	err := db.Ping()
	if err != nil {
		log.Println(">>> db error")
		panic(err)
	}
	defer db.Close()
	log.Println(">>> init database")

	err = infra.CreateTable(db)
	if err != nil {
		log.Fatalf(err.Error())
		panic(err)
	}
	log.Println(">>> creating table")

	// init repositories
	userRepo := repository.NewUserRepo(db)
	err = userRepo.InitUser()
	if err != nil {
		log.Fatalf(err.Error())
	}
	barangRepo := repository.NewBarangRepo(db)
	barangKeluarRepo := repository.NewBarangKeluarRepo(db)
	rkbmdRepo := repository.NewRkbmdRepo(db)

	// init services
	userService := app.NewUserService(userRepo)
	jurnalService := app.NewJurnalService(barangRepo)

	// init handlers
	authHandler := handler.NewAuthHandler(userService)
	jurnalHandler := handler.NewJurnalHandler(jurnalService)
	brgHandler := handler.NewBarangHandler(barangRepo)
	brgKeluarHandler := handler.NewBarangKeluarHandler(barangRepo, barangKeluarRepo)
	rkbmdHandler := handler.NewRkbmdHandler(rkbmdRepo)

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

	// auth
	router.HandleFunc("/api/auth/check", authHandler.Authorize).Methods("POST")

	// users
	router.HandleFunc("/api/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/api/users", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users", handler.CreateUser).Methods("POST")

	// barang
	router.HandleFunc("/api/barang/{id}", brgHandler.Get).Methods("GET")
	router.HandleFunc("/api/barang", brgHandler.Update).Methods("PUT")
	router.HandleFunc("/api/barang", brgHandler.GetAll).Methods("GET")
	router.HandleFunc("/api/barang", brgHandler.Save).Methods("POST")
	router.HandleFunc("/api/barang-keluar", brgKeluarHandler.Save).Methods("POST")
	router.HandleFunc("/api/barang-keluar", brgKeluarHandler.GetAll).Methods("GET")

	// jurnal
	router.HandleFunc("/api/jurnal", jurnalHandler.GetJurnal).Methods("GET")

	// rkbmd
	router.HandleFunc("/api/rkbmd", rkbmdHandler.Save).Methods("POST")
	router.HandleFunc("/api/rkbmd", rkbmdHandler.GetAll).Methods("GET")
	router.HandleFunc("/api/rkbmd", rkbmdHandler.Update).Methods("PUT")

	// frontend
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir + "/")))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)

	log.Printf("Starting server on http://localhost:%s ...", port)
	log.Fatal(http.ListenAndServe(":"+port, n))
}
