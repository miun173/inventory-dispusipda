package repository

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	"github.com/miun173/inventory-dispusibda/inventory/models"
	"github.com/miun173/inventory-dispusibda/inventory/models/rkbmd"
	"github.com/pkg/errors"
)

// RkbmdRepo repo of barang
type RkbmdRepo interface {
	GetAll() ([]models.RkbmdDetail, error)
	Save(brg *models.Rkbmd) error
	SaveDetail(dRkbmd []models.DetailRkbmd, rkbmdID int) error
	UpdateDetail(rd models.RkbmdDetail) error
}

type rkbmdRepo struct {
	db *sql.DB
}

// NewRkbmdRepo get rkbmd repo
func NewRkbmdRepo(db *sql.DB) RkbmdRepo {
	return &rkbmdRepo{db}
}

// Save repo
func (r *rkbmdRepo) Save(rkbdm *models.Rkbmd) error {
	q := "INSERT INTO rkbmd (tglBuat, status) VALUES(?, ?)"
	stm, err := r.db.Prepare(q)
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

// SaveDetail repo
func (r *rkbmdRepo) SaveDetail(dRkbmd []models.DetailRkbmd, rkbmdID int) error {
	var q bytes.Buffer
	q.WriteString("INSERT INTO detailRkbmd (rkbmdID, jml, namaBarang, status, kodeBarang, harga) VALUES")
	for i, d := range dRkbmd {
		var s string
		// if it's last element
		if i == len(dRkbmd)-1 {
			s = fmt.Sprintf("(%d, %d, '%s', '%s', '%s', %f);", rkbmdID, d.Jml, d.NamaBarang, rkbmd.RkbmdPending, d.KodeBarang, d.Harga)
		} else {
			s = fmt.Sprintf("(%d, %d, '%s', '%s', '%s', %f),", rkbmdID, d.Jml, d.NamaBarang, rkbmd.RkbmdPending, d.KodeBarang, d.Harga)
		}
		q.WriteString(s)
	}

	log.Println(q.String())

	rows, err := r.db.Query(q.String())
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

// GetAll repo
func (r *rkbmdRepo) GetAll() ([]models.RkbmdDetail, error) {
	rkbmds := make([]models.RkbmdDetail, 0)
	q := "SELECT id, tglBuat, status FROM rkbmd"
	rows, err := r.db.Query(q)
	if err != nil {
		return rkbmds, errors.WithStack(err)
	}
	defer rows.Close()

	var rd models.RkbmdDetail
	for rows.Next() {
		var status sql.NullString
		if err := rows.Scan(&rd.ID, &rd.TglBuat, &status); err != nil {
			return rkbmds, errors.WithStack(err)
		}

		if status.Valid {
			rd.Status = status.String
		}

		q = "SELECT id, rkbmdID, jml, namaBarang, status, kodeBarang, harga FROM detailRkbmd WHERE rkbmdID = ?"
		rows2, err := r.db.Query(q, rd.ID)
		if err != nil {
			return rkbmds, errors.WithStack(err)
		}

		var d models.DetailRkbmd
		rd.Detail = make([]models.DetailRkbmd, 0)
		for rows2.Next() {
			if err = rows2.Scan(&d.ID, &d.RkbmdID, &d.Jml, &d.NamaBarang, &d.Status, &d.KodeBarang, &d.Harga); err != nil {
				return rkbmds, errors.WithStack(err)
			}

			d.HargaTotal = d.Harga * float64(d.Jml)
			rd.Detail = append(rd.Detail, d)
		}
		rows2.Close()

		rkbmds = append(rkbmds, rd)
	}

	return rkbmds, nil
}

// UpdateDetail repo
func (r *rkbmdRepo) UpdateDetail(rd models.RkbmdDetail) error {
	q := "UPDATE rkbmd SET status = ? WHERE id = ?"
	stmt, err := r.db.Prepare(q)
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
		stmt, err = r.db.Prepare(q)
		if err != nil {
			return errors.WithStack(err)
		}
		_, err = stmt.Exec(d.Status, d.ID)
		stmt.Close()
	}

	return nil
}
