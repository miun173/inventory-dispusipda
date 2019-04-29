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

var people []models.Person

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	params := mux.Vars(r)

	personID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	for i, p := range people {
		if p.ID == personID {
			people[i] = person
			json.NewEncoder(w).Encode(person)
			break
		}
	}
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people := repo.GetAllPeople()
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	personID, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range repo.GetAllPeople() {
		if p.ID == personID {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	json.NewEncoder(w).Encode("Person not found")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	repo.Insert(person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// personID, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, p := range repo.GetAllPeople() {
	// 	if p.ID == personID {
	// 		repo.DeletePerson(p)
	// 		json.NewEncoder(w).Encode(p)
	// 		return
	// 	}
	// }

	json.NewEncoder(w).Encode("person not found")
}
