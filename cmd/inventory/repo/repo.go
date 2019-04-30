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

	stmts := []string{
		"CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, firstname TEXT, lastname TEXT, password TEXT, role TEXT);",
		"CREATE TABLE IF NOT EXISTS barang (id INTEGER PRIMARY KEY AUTOINCREMENT, kode TEXT, nama TEXT, reg TEXT, merk TEXT, ukuran TEXT, bahan TEXT, tglMasuk NUMERIC, tipeSpek TEXT, nomorSpek TEXT, caraPerolehan TEXT, jml INTEGER, harga REAL);",
		"CREATE TABLE IF NOT EXISTS barangKeluar (id INTEGER PRIMARY KEY AUTOINCREMENT, barangID INTEGER, jml INTEGER, tglKeluar NUMERIC);",
	}

	for _, s := range stmts {
		statement, err := db.Prepare(s)
		if err != nil {
			log.Fatal(err)
			return
		}
		statement.Exec()
	}

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

// CreateBarang insert new barang into db
func CreateBarang(brg *models.Barang) error {
	stm, err := db.Prepare(`INSERT INTO barang (kode, nama, reg, merk, ukuran, bahan, tglMasuk, tipeSpek, nomorSpek, caraPerolehan, harga) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	res, err := stm.Exec(brg.Kode, brg.Nama, brg.Reg, brg.Merk, brg.Ukuran, brg.Bahan, brg.TglMasuk, brg.TipeSpek, brg.NomorSpek, brg.CaraPerolehan, brg.Harga)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := res.LastInsertId()
	brg.ID = int(id)

	return nil
}

// CreateBarangKeluar insert new barang keluar to DB
func CreateBarangKeluar(brg *models.BarangKeluar) error {
	stm, err := db.Prepare(`INSERT INTO barangKeluar (barangID, jml, tglKeluar) 
		VALUES (?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
		return err
	}

	res, err := stm.Exec(brg.BarangID, brg.Jml, brg.TglKeluar)
	if err != nil {
		log.Fatal(err)
		return err
	}

	id, err := res.LastInsertId()
	brg.ID = int(id)

	return nil
}

// GetBarang get barang by its ID
func GetBarang(id int, brg *models.Barang) error {
	q := `SELECT 
		id, kode, nama, reg, merk, ukuran, bahan, tglMasuk, tipeSpek, nomorSpek, caraPerolehan, harga
		FROM barang
		WHERE id = ?`
	rows, err := db.Query(q, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&brg.ID, &brg.Kode, &brg.Nama, &brg.Reg, &brg.Merk, &brg.Ukuran, &brg.Bahan, &brg.TglMasuk, &brg.TipeSpek, &brg.NomorSpek, &brg.CaraPerolehan, &brg.Harga); err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

// GetAllBarang get all branangs in db
func GetAllBarang() ([]models.Barang, error) {
	q := `SELECT 
		id, kode, nama, reg, merk, ukuran, bahan, tglMasuk, tipeSpek, nomorSpek, caraPerolehan, harga
		FROM barang`

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var brgs []models.Barang
	var brg models.Barang
	for rows.Next() {
		if err := rows.Scan(&brg.ID, &brg.Kode, &brg.Nama, &brg.Reg, &brg.Merk, &brg.Ukuran, &brg.Bahan, &brg.TglMasuk, &brg.TipeSpek, &brg.NomorSpek, &brg.CaraPerolehan, &brg.Harga); err != nil {
			log.Fatal(err)
			return nil, err
		}

		brgs = append(brgs, brg)
	}

	return brgs, nil
}
