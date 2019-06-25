package repository

import (
	"database/sql"
	"fmt"
	"log"

	// sqlite driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
)

// BarangRepo repo of barang
type BarangRepo interface {
	GetAll() ([]models.Barang, error)
	Get(id int) (models.Barang, error)
	Save(brg *models.Barang) error
	IsExists(barangID int) (bool, error)
	Update(brg models.Barang) error
}

type barangRepo struct {
	db *sql.DB
}

// NewBarangRepo get barang repo
func NewBarangRepo(db *sql.DB) BarangRepo {
	return &barangRepo{db}
}

func (b *barangRepo) GetAll() ([]models.Barang, error) {
	q := `SELECT id,
			kode,
			nama,
			reg,
			merk,
			jml,
			ket,
			ukuran,
			bahan,
			tglmasuk,
			tipespek,
			nomorspek,
			caraperolehan,
			harga,
			nilaisisa,
			umurekonomis,
			umurpenggunaan,
			nilaibuku,
			bebanpenyusutan
		FROM   barang `

	rows, err := b.db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "select all barang error")
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	brgs := make([]models.Barang, 0)
	var brg models.Barang
	for rows.Next() {
		if err := rows.Scan(
			&brg.ID,
			&brg.Kode,
			&brg.Nama,
			&brg.Reg,
			&brg.Merk,
			&brg.Jml,
			&brg.Ket,
			&brg.Ukuran,
			&brg.Bahan,
			&brg.TglMasuk,
			&brg.TipeSpek,
			&brg.NomorSpek,
			&brg.CaraPerolehan,
			&brg.Harga,
			&brg.NilaiSisa,
			&brg.UmurEkonomis,
			&brg.UmurPenggunaan,
			&brg.NilaiBuku,
			&brg.BebanPenyusutan,
		); err != nil {
			err = errors.Wrap(err, "scan barang rows error")
			log.Printf("%+v\n", err)
			return nil, err
		}

		brgs = append(brgs, brg)
	}

	return brgs, nil
}

// Get get barang by its ID
func (b *barangRepo) Get(id int) (models.Barang, error) {
	var barang models.Barang
	q := `SELECT id,
			kode,
			nama,
			reg,
			merk,
			jml,
			ket,
			ukuran,
			bahan,
			tglmasuk,
			tipespek,
			nomorspek,
			caraperolehan,
			harga,
			nilaisisa,
			umurekonomis,
			umurpenggunaan,
			nilaibuku,
			bebanpenyusutan,
			koreksi
		FROM   barang
		WHERE  id = ? `
	rows, err := b.db.Query(q, id)
	if err != nil {
		err = errors.Wrap(err, "select barang error")
		log.Printf("%+v\n", err)
		return barang, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&barang.ID,
			&barang.Kode,
			&barang.Nama,
			&barang.Reg,
			&barang.Merk,
			&barang.Jml,
			&barang.Ket,
			&barang.Ukuran,
			&barang.Bahan,
			&barang.TglMasuk,
			&barang.TipeSpek,
			&barang.NomorSpek,
			&barang.CaraPerolehan,
			&barang.Harga,
			&barang.NilaiSisa,
			&barang.UmurEkonomis,
			&barang.UmurPenggunaan,
			&barang.NilaiBuku,
			&barang.BebanPenyusutan,
			&barang.Koreksi,
		); err != nil {
			err = errors.Wrap(err, "scan barang rows error")
			log.Printf("%+v\n", err)
			return barang, err
		}
	}

	return barang, nil
}

// Save insert new barang into db
func (b *barangRepo) Save(brg *models.Barang) error {
	stm, err := b.db.Prepare(`INSERT INTO barang
			(kode,
			nama,
			reg,
			merk,
			jml,
			ket,
			ukuran,
			bahan,
			tglmasuk,
			tipespek,
			nomorspek,
			caraperolehan,
			harga,
			nilaisisa,
			umurekonomis,
			umurpenggunaan,
			nilaibuku,
			bebanpenyusutan,
			koreksi)
		VALUES      (?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?) `)
	if err != nil {
		err = errors.Wrap(err, "insert barang error")
		log.Printf("%+v\n", err)
		return err
	}
	defer stm.Close()

	res, err := stm.Exec(
		brg.Kode,
		brg.Nama,
		brg.Reg,
		brg.Merk,
		brg.Jml,
		brg.Ket,
		brg.Ukuran,
		brg.Bahan,
		brg.TglMasuk,
		brg.TipeSpek,
		brg.NomorSpek,
		brg.CaraPerolehan,
		brg.Harga,
		brg.NilaiSisa,
		brg.UmurEkonomis,
		brg.UmurPenggunaan,
		brg.NilaiBuku,
		brg.BebanPenyusutan,
		brg.Koreksi,
	)
	if err != nil {
		err = errors.Wrap(err, "insert barang error")
		log.Printf("%+v\n", err)

		return err
	}

	id, err := res.LastInsertId()
	brg.ID = int(id)

	return nil
}

// IsExists check if barang exists by its ID
func (b *barangRepo) IsExists(barangID int) (bool, error) {
	var exists bool
	q := fmt.Sprintf("SELECT COUNT(1) FROM barang WHERE id = %d", barangID)
	err := b.db.QueryRow(q).Scan(&exists)
	if err != nil {
		err = errors.Wrap(err, "check barang exists error")
		log.Printf("%+v\n", err)
		return false, err
	}

	log.Printf("exists, %v", exists)

	return exists, nil
}

// Update repo
func (b *barangRepo) Update(brg models.Barang) error {
	stm, err := b.db.Prepare(`UPDATE barang 
	SET 
		kode = ?, nama = ?, reg = ?, merk = ?, jml = ?, ket = ?, ukuran = ?, bahan = ?, tglMasuk = ?, tipeSpek = ?, nomorSpek = ?, caraPerolehan = ?, harga = ?, nilaiSisa = ?, umurEkonomis = ?, umurPenggunaan = ?, nilaiBuku = ?, bebanPenyusutan = ?
	WHERE
		id = ?
	`)

	_, err = stm.Exec(brg.Kode, brg.Nama, brg.Reg, brg.Merk, brg.Jml, brg.Ket, brg.Ukuran, brg.Bahan, brg.TglMasuk, brg.TipeSpek, brg.NomorSpek, brg.CaraPerolehan, brg.Harga, brg.NilaiSisa, brg.UmurEkonomis, brg.UmurPenggunaan, brg.NilaiBuku, brg.BebanPenyusutan, brg.ID)
	if err != nil {
		return err
	}

	return nil
}
