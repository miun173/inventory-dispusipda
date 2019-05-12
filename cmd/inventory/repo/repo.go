package repo

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/models/rkbmd"
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
		"CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, firstname TEXT, lastname TEXT, password TEXT, role TEXT, token TEXT);",
		"CREATE TABLE IF NOT EXISTS rkbmd (id INTEGER PRIMARY KEY AUTOINCREMENT, tglBuat NUMERIC, status TEXT);",
		"CREATE TABLE IF NOT EXISTS detailRkbmd (id INTEGER PRIMARY KEY AUTOINCREMENT, rkbmdID INTEGER, namaBarang TEXT, kodeBarang TEXT, jml NUMERIC, status TEXT, FOREIGN KEY (rkbmdID) REFERENCES rkbmd(id))",
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

	initUser()

	log.Println("connect to db")
}

func initUser() {
	users := [3][2]string{{"divisi", "divisi"}, {"subagumum", "pejabat"}, {"petugasBarang", "petugasBarang"}}
	for _, u := range users {
		user := models.User{
			Username:  u[0],
			FirstName: u[0],
			LastName:  u[0],
			Role:      u[1],
			Password:  "123",
		}

		exists, err := CheckUsernameExist(user)
		if err != nil {
			log.Fatalf("%+v\n", err)
			return
		}

		if !exists {
			err = CreateUser(&user)
			if err != nil {
				log.Fatalf("%+v\n", err)
				return
			}
		}
	}
}

// CheckUsernameExist repo
func CheckUsernameExist(user models.User) (bool, error) {
	var exists bool
	q := fmt.Sprintf("SELECT COUNT(1) FROM users WHERE username = '%s'", user.Username)
	err := db.QueryRow(q).Scan(&exists)
	if err != nil {
		err = errors.WithStack(err)
		return false, err
	}

	return exists, nil
}

// CreateUser insert new user to db
func CreateUser(user *models.User) error {
	stm, err := db.Prepare("INSERT INTO users (username, firstname, lastname, password, role, token) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		err = errors.Wrap(err, "create user error")
		log.Printf("%+v\n", err)
		return err
	}

	secret := time.Now().String()
	token := base64.StdEncoding.EncodeToString([]byte(string(secret)))

	res, err := stm.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Role, token)
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

// GetUserByID get user by id
func GetUserByID(user *models.User, userID int) error {
	q := "SELECT id, firstname, lastname, password, role, token FROM users WHERE id = ? "
	err := db.QueryRow(q, userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.Token)
	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "select user error")
		log.Printf("%+v\n", err)
		return err
	}

	return nil
}

// GetUserByUsername find user by given username
func GetUserByUsername(user *models.User, username string) error {
	q := "SELECT id, firstname, lastname, password, role, token FROM users WHERE username= ? "
	err := db.QueryRow(q, username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.Token)

	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "select user error")
		log.Printf("%+v\n", err)
		return err
	}

	return nil
}

// GetAllUser query all users
func GetAllUser() ([]models.User, error) {
	users := make([]models.User, 0)
	q := "SELECT id, username, firstname, lastname, role, token FROM users"
	rows, err := db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "select all user error")
		log.Printf("%+v\n", err)
		return users, err
	}
	defer rows.Close()

	var u models.User
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Role, &u.Token); err != nil {
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

// CreateRkbmd repo
func CreateRkbmd(rkbdm *models.Rkbmd) error {
	q := "INSERT INTO rkbmd (tglBuat, status) VALUES(?, ?)"
	stm, err := db.Prepare(q)
	if err != nil {
		return errors.WithStack(err)
	}

	res, err := stm.Exec(rkbdm.TglBuat, rkbmd.RkbmdPending)
	if err != nil {
		return errors.WithStack(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.WithStack(err)
	}

	rkbdm.ID = int(id)
	return nil
}

// CreateDetailRkbmd repo
func CreateDetailRkbmd(dRkbmd []models.DetailRkbmd, rkbmdID int) error {
	var q bytes.Buffer
	q.WriteString("INSERT INTO detailRkbmd (rkbmdID, jml, namaBarang, status, kodeBarang) VALUES")
	for i, d := range dRkbmd {
		var s string
		// if it's last element
		if i == len(dRkbmd)-1 {
			s = fmt.Sprintf("(%d, %d, '%s', '%s', '%s');", rkbmdID, d.Jml, d.NamaBarang, rkbmd.RkbmdPending, d.KodeBarang)
		} else {
			s = fmt.Sprintf("(%d, %d, '%s', '%s', '%s'),", rkbmdID, d.Jml, d.NamaBarang, rkbmd.RkbmdPending, d.KodeBarang)
		}
		q.WriteString(s)
	}

	rows, err := db.Query(q.String())
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	var ds []models.DetailRkbmd
	var d models.DetailRkbmd
	for rows.Next() {
		if err := rows.Scan(&d.ID); err != nil {
			return errors.WithStack(err)
		}

		log.Println(d)
		ds = append(ds, d)
	}

	return nil
}

// GetAllRkbmd repo
func GetAllRkbmd() ([]models.RkbmdDetail, error) {
	rkbmds := make([]models.RkbmdDetail, 0)
	q := "SELECT id, tglBuat, status FROM rkbmd"
	rows, err := db.Query(q)
	if err != nil {
		return rkbmds, errors.WithStack(err)
	}
	defer rows.Close()

	var r models.RkbmdDetail
	for rows.Next() {
		var status sql.NullString
		if err := rows.Scan(&r.ID, &r.TglBuat, &status); err != nil {
			return rkbmds, errors.WithStack(err)
		}

		if status.Valid {
			r.Status = status.String
		}

		q = "SELECT id, rkbmdID, jml, namaBarang, status, kodeBarang FROM detailRkbmd WHERE rkbmdID = ?"
		rows2, err := db.Query(q, r.ID)
		if err != nil {
			return rkbmds, errors.WithStack(err)
		}

		var d models.DetailRkbmd
		r.Detail = make([]models.DetailRkbmd, 0)
		for rows2.Next() {
			if err = rows2.Scan(&d.ID, &d.RkbmdID, &d.Jml, &d.NamaBarang, &d.Status, &d.KodeBarang); err != nil {
				return rkbmds, errors.WithStack(err)
			}

			r.Detail = append(r.Detail, d)
		}
		rows2.Close()

		rkbmds = append(rkbmds, r)
	}

	return rkbmds, nil
}

// UpdateRkbmdDetail repo
func UpdateRkbmdDetail(rd models.RkbmdDetail) error {
	q := "UPDATE rkbmd SET status = ? WHERE id = ?"
	stmt, err := db.Prepare(q)
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = stmt.Exec(rd.Status, rd.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	stmt.Close()

	for _, d := range rd.Detail {
		q = "UPDATE detailRkbmd SET status = ? WHERE id = ?"
		stmt, err = db.Prepare(q)
		if err != nil {
			return errors.WithStack(err)
		}
		_, err = stmt.Exec(d.Status, d.ID)
		stmt.Close()
	}

	return nil
}
