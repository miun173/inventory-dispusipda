package repository

import (
	"database/sql"
	"log"

	"github.com/miun173/inventory-dispusibda/inventory/models"
	"github.com/pkg/errors"
)

// BarangKeluarRepo repo of barang
type BarangKeluarRepo interface {
	GetAll() ([]models.BarangKeluar, error)
	Save(brg *models.BarangKeluar) error
}

type barangKeluarRepo struct {
	db *sql.DB
}

// NewBarangKeluarRepo get barang repo
func NewBarangKeluarRepo(db *sql.DB) BarangKeluarRepo {
	return &barangKeluarRepo{db}
}

// GetAll get all barang keluar in db
func (b *barangKeluarRepo) GetAll() ([]models.BarangKeluar, error) {
	brgs := make([]models.BarangKeluar, 0)

	q := `SELECT bk.id, bk.barangID, bk.jml, bk.tglKeluar, b.nama
	FROM barangKeluar bk, barang b
	WHERE bk.barangID=b.id`

	rows, err := b.db.Query(q)
	if err != nil {
		err = errors.Wrap(err, "get all barang keluar error")
		log.Printf("%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	var bk models.BarangKeluar
	for rows.Next() {
		if err := rows.Scan(&bk.ID, &bk.BarangID, &bk.Jml, &bk.TglKeluar, &bk.Nama); err != nil {
			err = errors.Wrap(err, "scan barang keluar error")
			log.Printf("%+v\n", err)
			return nil, err
		}

		brgs = append(brgs, bk)
	}

	return brgs, nil
}

// Save insert new barang keluar to DB
func (b *barangKeluarRepo) Save(brg *models.BarangKeluar) error {
	stm, err := b.db.Prepare(`INSERT INTO barangKeluar (barangID, jml, tglKeluar) 
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
