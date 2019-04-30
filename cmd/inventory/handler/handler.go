package handler

import (
	"encoding/json"
	"log"
	"net/http"
	// "strconv"
	// "github.com/gorilla/mux"
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

// func UpdatePerson(w http.ResponseWriter, r *http.Request) {
// 	var person models.Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	params := mux.Vars(r)

// 	personID, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i, p := range people {
// 		if p.ID == personID {
// 			people[i] = person
// 			json.NewEncoder(w).Encode(person)
// 			break
// 		}
// 	}
// }

// func GetPeople(w http.ResponseWriter, r *http.Request) {
// 	people := repo.GetAllPeople()
// 	json.NewEncoder(w).Encode(people)
// }

// func GetPerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	personID, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, p := range repo.GetAllPeople() {
// 		if p.ID == personID {
// 			json.NewEncoder(w).Encode(p)
// 			return
// 		}
// 	}

// 	json.NewEncoder(w).Encode("Person not found")
// }

// func CreatePerson(w http.ResponseWriter, r *http.Request) {
// 	var person models.Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	repo.Insert(person)
// 	json.NewEncoder(w).Encode(person)
// }

// func DeletePerson(w http.ResponseWriter, r *http.Request) {
// 	// params := mux.Vars(r)
// 	// personID, err := strconv.Atoi(params["id"])
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// for _, p := range repo.GetAllPeople() {
// 	// 	if p.ID == personID {
// 	// 		repo.DeletePerson(p)
// 	// 		json.NewEncoder(w).Encode(p)
// 	// 		return
// 	// 	}
// 	// }

// 	json.NewEncoder(w).Encode("person not found")
// }
