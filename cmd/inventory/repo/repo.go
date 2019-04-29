package repo

import (
	"database/sql"
	"log"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
)

var db *sql.DB

// InitDB initialize connection & creating database
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, firstname TEXT, lastname TEXT, password TEXT);
		CREATE TABLE IF NOT EXISTS barang (id INTEGER PRIMARY KEY AUTOINCREMENT, kode TEXT, nama TEXT, reg TEXT, merk TEXT, ukuran TEXT, bahan TEXT, tglMasuk NUMERIC, tipeSpek TEXT, nomorSpek INTEGER, caraPerolehan TEXT, harga REAL);
		CREATE TABLE IF NOT EXISTS barangKeluar (id INTEGER PRIMARY KEY AUTOINCREMENT, idBarang INTEGER, jumlah INTEGER, tglKeluar NUMERIC);
	`)
	if err != nil {
		log.Fatal(err)
		return
	}
	statement.Exec()

	log.Println("connect to db")
}

// CreateUser insert new user to db
func CreateUser(user *models.User) error {
	stm, err := db.Prepare("INSERT INTO users (firstname, lastname, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}

	res, err := stm.Exec(user.FirstName, user.LastName, user.Password)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := res.LastInsertId()
	user.ID = int(id)
	return nil
}

// GetAllUser query all users
func GetAllUser() ([]models.User, error) {
	var users []models.User
	q := "SELECT id, firstname, lastname FROM users"
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
		return users, err
	}

	var u models.User
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName); err != nil {
			log.Fatal(err)
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}
