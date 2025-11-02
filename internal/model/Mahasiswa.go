package model

type Mahasiswa struct {
	Nim            int64  `json:"nim"`
	Nama_mahasiswa string `json:"nama"`
	Tempat_lahir   string `json:"tempat_lahir"`
	Tanggal_lahir  string `json:"tanggal_lahir"`
	Alamat_kos     string `json:"alamat_kos"`
}
