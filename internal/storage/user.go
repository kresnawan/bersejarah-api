package storage

import (
	"app/internal/model"
	"encoding/json"
	"log"
)

func DbConnect() []byte {

	var mahasiswas []model.Mahasiswa
	rows, err := Db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var mhsw model.Mahasiswa
		if err := rows.Scan(&mhsw.Nim, &mhsw.Nama_mahasiswa, &mhsw.Tempat_lahir, &mhsw.Tanggal_lahir, &mhsw.Alamat_kos); err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, mhsw)
	}

	jsonData, err := json.Marshal(mahasiswas)
	if err != nil {
		log.Fatal(err)
	}
	return jsonData
}
