package models

// Person model
type Person struct {
	ID        int    `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

// User model
type User struct {
	ID        int    `json:"id,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

// Barang model
type Barang struct {
	ID              int     `json:"id"`
	Kode            string  `json:"kode"`
	Nama            string  `json:"nama"`
	Reg             string  `json:"reg"`
	Merk            string  `json:"merk"`
	Jml             int     `json:"jml"`
	Ukuran          string  `json:"ukuran"`
	Bahan           string  `json:"bahan"`
	TipeSpek        string  `json:"tipeSpek"`
	NomorSpek       string  `json:"nomorSpek"`
	CaraPerolehan   string  `json:"caraPerolehan"`
	TglMasuk        int     `json:"tglMasuk"`
	Harga           float64 `json:"harga"`
	NilaiSisa       float64 `json:"nilaiSisa"`
	UmurEkonomis    int     `json:"umurEkonomis"`
	UmurPenggunaan  int     `json:"umurPenggunaan"`
	NilaiBuku       float64 `json:"nilaiBuku"`
	BebanPenyusutan float64 `json:"bebanPenyusutan"`
	Koreksi         float64 `json:"koreksi,omitempty"`
	Ket             string  `json:"ket"`
}

// BarangKeluar model
type BarangKeluar struct {
	Nama      string `json:"nama,omitempty"`
	ID        int    `json:"id"`
	BarangID  int    `json:"barangID"`
	Jml       int    `json:"jml"`
	TglKeluar int    `json:"tglKeluar"`
}

// Rkbmd model
type Rkbmd struct {
	ID      int    `json:"id"`
	TglBuat int    `json:"tglBuat"`
	Status  string `json:"status"`
}

// DetailRkbmd model
type DetailRkbmd struct {
	ID         int    `json:"id,omitempty"`
	RkbmdID    int    `json:"rkbmdID,omitempty"`
	Jml        int    `json:"jml,omitempty"`
	KodeBarang string `json:"kodeBarang,omitempty"`
	NamaBarang string `json:"namaBarang,omitempty"`
	Status     string `json:"status,omitempty"`
}

// RkbmdDetail model
type RkbmdDetail struct {
	Rkbmd
	Detail []DetailRkbmd `json:"details"`
}

// Jurnal model
type Jurnal struct {
	Barang
	Penyusutan float64 `json:"penyusutan"`
}
