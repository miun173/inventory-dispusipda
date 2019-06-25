package infra

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// NewDB initialize connection & creating database
func NewDB() *sql.DB {
	var conn *sql.DB
	conn, err := sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Printf("%+v\n", errors.WithStack(err))
		return conn
	}

	return conn
}

// CreateTable creat a table
func CreateTable(db *sql.DB) error {
	stmts := []string{
		"CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, firstname TEXT, lastname TEXT, password TEXT, role TEXT, token TEXT);",
		"CREATE TABLE IF NOT EXISTS rkbmd (id INTEGER PRIMARY KEY AUTOINCREMENT, tglBuat NUMERIC, status TEXT);",
		"CREATE TABLE IF NOT EXISTS detailRkbmd (id INTEGER PRIMARY KEY AUTOINCREMENT, rkbmdID INTEGER, namaBarang TEXT, kodeBarang TEXT, jml NUMERIC, status TEXT, harga REAL, FOREIGN KEY (rkbmdID) REFERENCES rkbmd(id))",
		"CREATE TABLE IF NOT EXISTS barang (id INTEGER PRIMARY KEY AUTOINCREMENT, kode TEXT, nama TEXT, reg TEXT, merk TEXT, ukuran TEXT, bahan TEXT, tglMasuk NUMERIC, tipeSpek TEXT, nomorSpek TEXT, caraPerolehan TEXT, jml INTEGER, ket TEXT, harga REAL, nilaiSisa REAL, umurEkonomis INTEGER, umurPenggunaan INTEGER, nilaiBuku REAL, bebanPenyusutan REAL, koreksi REAL);",
		"CREATE TABLE IF NOT EXISTS barangKeluar (id INTEGER PRIMARY KEY AUTOINCREMENT, barangID INTEGER, jml INTEGER, tglKeluar NUMERIC, FOREIGN KEY (barangID) REFERENCES barang(id));",
	}

	for _, s := range stmts {
		statement, err := db.Prepare(s)
		if err != nil {
			log.Printf("%+v\n", errors.WithStack(err))
			return err
		}
		statement.Exec()
		statement.Close()
	}

	return nil
}
