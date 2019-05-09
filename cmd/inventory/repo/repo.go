package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"

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
		log.Printf("%+v\n", errors.WithStack(err))
		return
	}

	stmts := []string{
		"CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, firstname TEXT, lastname TEXT, password TEXT, role TEXT);",
		"CREATE TABLE IF NOT EXISTS rkbmd (id INTEGER PRIMARY KEY AUTOINCREMENT, tglBuat NUMERIC);",
		"CREATE TABLE IF NOT EXISTS detailRkbmd (id INTEGER PRIMARY KEY AUTOINCREMENT, rkbmdID INTEGER, namaBarang TEXT, jml NUMERIC, status TEXT, FOREIGN KEY (rkbmdID) REFERENCES rkbmd(id))",
		"CREATE TABLE IF NOT EXISTS barang (id INTEGER PRIMARY KEY AUTOINCREMENT, kode TEXT, nama TEXT, reg TEXT, merk TEXT, ukuran TEXT, bahan TEXT, tglMasuk NUMERIC, tipeSpek TEXT, nomorSpek TEXT, caraPerolehan TEXT, jml INTEGER, ket TEXT, harga REAL, nilaiSisa REAL, umurEkonomis INTEGER, umurPenggunaan INTEGER, nilaiBuku REAL, bebanPenyusutan REAL, koreksi REAL);",
		"CREATE TABLE IF NOT EXISTS barangKeluar (id INTEGER PRIMARY KEY AUTOINCREMENT, barangID INTEGER, jml INTEGER, tglKeluar NUMERIC, FOREIGN KEY (barangID) REFERENCES barang(id));",
	}

	for _, s := range stmts {
		statement, err := db.Prepare(s)
		if err != nil {
			log.Printf("%+v\n", errors.WithStack(err))
			return
		}
		statement.Exec()
		statement.Close()
	}

	log.Println("connect to db")
}

// CreateUser insert new user to db
func CreateUser(user *models.User) error {
	stm, err := db.Prepare("INSERT INTO users (firstname, lastname, password, role) VALUES (?, ?, ?, ?)")
	if err != nil {
		err = errors.Wrap(err, "create user error")
		log.Printf("%+v\n", err)
		return err
	}

	res, err := stm.Exec(user.FirstName, user.LastName, user.Password, user.Role)
	if err != nil {
		err = errors.Wrap(err, "create user error")
		log.Printf("%+v\n", err)
		return err
	}
	defer stm.Close()

	id, err := res.LastInsertId()
	user.ID = int(id)
	return nil
}

// GetUser find user by given username
func GetUser(user *models.User) error {
	q := "SELECT id, firstname, lastname, password, role FROM users WHERE username= ? "
	err := db.QueryRow(q, user.Username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Role)
	if err != nil {
		err = errors.Wrap(err, "select user error")
		log.Printf("%+v\n", err)
		return err
	}

	return nil
}

// GetAllUser query all users
func GetAllUser() ([]models.User, error) {
	users := make([]models.User, 0)
	q := "SELECT id, firstname, lastname, role FROM users"
	rows, err := db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "select all user error")
		log.Printf("%+v\n", err)
		return users, err
	}
	defer rows.Close()

	var u models.User
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Role); err != nil {
			err = errors.Wrap(err, "scan users row error")
			log.Printf("%+v\n", err)
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

// CreateBarang insert new barang into db
func CreateBarang(brg *models.Barang) error {
	stm, err := db.Prepare(`INSERT INTO barang (kode, nama, reg, merk, jml, ket, ukuran, bahan, tglMasuk, tipeSpek, nomorSpek, caraPerolehan, harga, nilaiSisa, umurEkonomis, umurPenggunaan, nilaiBuku, bebanPenyusutan, koreksi) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		err = errors.Wrap(err, "insert barang error")
		log.Printf("%+v\n", err)
		return err
	}
	defer stm.Close()

	res, err := stm.Exec(brg.Kode, brg.Nama, brg.Reg, brg.Merk, brg.Jml, brg.Ket, brg.Ukuran, brg.Bahan, brg.TglMasuk, brg.TipeSpek, brg.NomorSpek, brg.CaraPerolehan, brg.Harga, brg.NilaiSisa, brg.UmurEkonomis, brg.UmurPenggunaan, brg.NilaiBuku, brg.BebanPenyusutan, brg.Koreksi)
	if err != nil {
		err = errors.Wrap(err, "insert barang error")
		log.Printf("%+v\n", err)

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
		err = errors.Wrap(err, "insert barang keluar error")
		log.Printf("%+v\n", err)
		return err
	}
	defer stm.Close()

	res, err := stm.Exec(brg.BarangID, brg.Jml, brg.TglKeluar)
	if err != nil {
		err = errors.Wrap(err, "insert barang keluar error")
		log.Printf("%+v\n", err)
		return err
	}

	id, err := res.LastInsertId()
	brg.ID = int(id)

	return nil
}

// GetBarang get barang by its ID
func GetBarang(id int, brg *models.Barang) error {
	q := `SELECT 
		id, kode, nama, reg, merk, jml, ket, ukuran, bahan, tglMasuk, tipeSpek, nomorSpek, caraPerolehan, harga, nilaiSisa, umurEkonomis, umurPenggunaan, nilaiBuku, bebanPenyusutan, koreksi
		FROM barang
		WHERE id = ?`
	rows, err := db.Query(q, id)
	if err != nil {
		err = errors.Wrap(err, "select barang error")
		log.Printf("%+v\n", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&brg.ID, &brg.Kode, &brg.Nama, &brg.Reg, &brg.Merk, &brg.Jml, &brg.Ket, &brg.Ukuran, &brg.Bahan, &brg.TglMasuk, &brg.TipeSpek, &brg.NomorSpek, &brg.CaraPerolehan, &brg.Harga, &brg.NilaiSisa, &brg.UmurEkonomis, &brg.UmurPenggunaan, &brg.NilaiBuku, &brg.BebanPenyusutan, &brg.Koreksi); err != nil {
			err = errors.Wrap(err, "scan barang rows error")
			log.Printf("%+v\n", err)
			return err
		}
	}

	return nil
}

// GetAllBarang get all barangs in db
func GetAllBarang() ([]models.Barang, error) {
	q := `SELECT 
		id, kode, nama, reg, merk, jml, ket, ukuran, bahan, tglMasuk, tipeSpek, nomorSpek, caraPerolehan, harga, nilaiSisa, umurEkonomis, umurPenggunaan, nilaiBuku, bebanPenyusutan
		FROM barang`

	rows, err := db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "select all barang error")
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	brgs := make([]models.Barang, 0)
	var brg models.Barang
	for rows.Next() {
		if err := rows.Scan(&brg.ID, &brg.Kode, &brg.Nama, &brg.Reg, &brg.Merk, &brg.Jml, &brg.Ket, &brg.Ukuran, &brg.Bahan, &brg.TglMasuk, &brg.TipeSpek, &brg.NomorSpek, &brg.CaraPerolehan, &brg.Harga, &brg.NilaiSisa, &brg.UmurEkonomis, &brg.UmurPenggunaan, &brg.NilaiBuku, &brg.BebanPenyusutan); err != nil {
			err = errors.Wrap(err, "scan barang rows error")
			log.Printf("%+v\n", err)
			return nil, err
		}

		brgs = append(brgs, brg)
	}

	return brgs, nil
}

// CheckBarangExists check if barang exists by its ID
func CheckBarangExists(barangID int) (bool, error) {
	var exists bool
	q := fmt.Sprintf("SELECT COUNT(1) FROM barang WHERE id = %d", barangID)
	err := db.QueryRow(q).Scan(&exists)
	if err != nil {
		err = errors.Wrap(err, "check barang exists error")
		log.Printf("%+v\n", err)
		return false, err
	}

	log.Printf("exists, %v", exists)

	return exists, nil
}

// GetAllBarangKeluar get all barnag keluar in db
func GetAllBarangKeluar() ([]models.BarangKeluar, error) {
	brgs := make([]models.BarangKeluar, 0)

	q := `SELECT bk.id, bk.barangID, bk.jml, bk.tglKeluar, b.nama
	FROM barangKeluar bk, barang b
	WHERE bk.barangID=b.id`

	rows, err := db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "get all barang keluar error")
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	var b models.BarangKeluar
	for rows.Next() {
		if err := rows.Scan(&b.ID, &b.BarangID, &b.Jml, &b.TglKeluar, &b.Nama); err != nil {
			err = errors.Wrap(err, "scan barang keluar error")
			log.Printf("%+v\n", err)
			return nil, err
		}

		brgs = append(brgs, b)
	}

	return brgs, nil
}
