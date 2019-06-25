package app

import (
	"time"

	"github.com/miun173/inventory-dispusibda/cmd/inventory/models"
	"github.com/miun173/inventory-dispusibda/cmd/inventory/repository"
)

// JurnalService jurnal domain service
type JurnalService interface {
	GetAll() ([]models.Jurnal, error)
}

type jurnalService struct {
	barangRepo repository.BarangRepo
}

// NewJurnalService get jurnal service object
func NewJurnalService(barangRepo repository.BarangRepo) JurnalService {
	return &jurnalService{barangRepo}
}

func (j *jurnalService) GetAll() ([]models.Jurnal, error) {
	jurnals := make([]models.Jurnal, 0)

	brgs, err := j.barangRepo.GetAll()
	if err != nil {
		return jurnals, err
	}

	for _, b := range brgs {
		b.BebanPenyusutan = (b.Harga - b.NilaiSisa) / float64(b.UmurEkonomis)

		tahunMasuk := time.Unix(int64(b.TglMasuk/1000), 0).Year()
		tahunBerjalan := time.Now().Year() - tahunMasuk

		penyusutanPerTahun := float64(tahunBerjalan) * b.BebanPenyusutan
		b.NilaiBuku = b.Harga - penyusutanPerTahun
		b.UmurPenggunaan = tahunBerjalan
		b.HargaTotal = b.Harga * float64(b.Jml)

		j := models.Jurnal{
			Barang:     b,
			Penyusutan: penyusutanPerTahun,
		}
		jurnals = append(jurnals, j)
	}

	return jurnals, nil
}
