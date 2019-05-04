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
}

// Barang model
type Barang struct {
	ID              int     `json:"id,omitempty"`
	Kode            string  `json:"kode,omitempty"`
	Nama            string  `json:"nama,omitempty"`
	Reg             string  `json:"reg,omitempty"`
	Merk            string  `json:"merk,omitempty"`
	Ukuran          string  `json:"ukuran,omitempty"`
	Bahan           string  `json:"bahan,omitempty"`
	TipeSpek        string  `json:"tipeSpek,omitempty"`
	NomorSpek       string  `json:"nomorSpek,omitempty"`
	CaraPerolehan   string  `json:"caraPerolehan,omitempty"`
	TglMasuk        int     `json:"tglMasuk,omitempty"`
	Harga           float64 `json:"harga,omitempty"`
	NilaiSisa       float64 `json:"nilaiSisa,omitempty"`
	UmurEkonomis    int     `json:"umurEkonomis,omitempty"`
	UmurPenggunaan  int     `json:"umurPenggunaan,omitempty"`
	NilaiBuku       float64 `json:"nilaiBuku,omitempty"`
	BebanPenyusutan float64 `json:"bebanPenyusutan,omitempty"`
}

// BarangKeluar model
type BarangKeluar struct {
	ID        int `json:"id,omitempty"`
	BarangID  int `json:"barangID,omitempty"`
	Jml       int `json:"jml,omitempty"`
	TglKeluar int `json:"tglKeluar,omitempty"`
}

// Jurnal model
type Jurnal struct {
	Barang
	Penyusutan float64 `json:"penyusutan,omitempty"`
}
