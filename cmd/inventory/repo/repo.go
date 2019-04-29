package repo

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./golang-rest.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS person (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err)
		return
	}
	statement.Exec()

	log.Println("connect to db")
}

func GetAllPeople() []models.Person {
	var people []models.Person
	q := "SELECT id, firstname, lastname FROM person"
	if db == nil {
		log.Fatal("db nil")
	}
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
		return people
	}

	var p models.Person
	for rows.Next() {
		rows.Scan(&p.ID, &p.Firstname, &p.Lastname)
		people = append(people, p)
	}

	return people
}

func Insert(person models.Person) {
	stm, err := db.Prepare("INSERT INTO person (firstname, lastname) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}

	stm.Exec(person.Firstname, person.Lastname)
}
