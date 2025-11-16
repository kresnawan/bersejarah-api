package storage

import (
	"app/internal/model"
	"encoding/json"
)

func InsertDataTempat(namaTempat, alamatTempat, deskripsiTempat string, lati, longi float32) (int64, error) {
	insertResult, err := Db.Exec("INSERT INTO data_tempat (title, description, address, latitude, longitude) VALUES (?, ?, ?, ?, ?)", namaTempat, deskripsiTempat, alamatTempat, lati, longi)

	if err != nil {
		return 0, err
	}

	insertId, err := insertResult.LastInsertId()

	if err != nil {
		return 0, err
	}

	return insertId, nil
}

func GetAllDataTempat() ([]byte, error) {
	type DataTempat struct {
		Id          int64   `json:"id"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Address     string  `json:"address"`
	}
	var dataArr []DataTempat
	rows, err := Db.Query("SELECT id, latitude, longitude, title, description, address FROM data_tempat")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var data DataTempat
		if err := rows.Scan(&data.Id, &data.Latitude, &data.Longitude, &data.Title, &data.Description, &data.Address); err != nil {
			return nil, err
		}

		dataArr = append(dataArr, data)
	}

	jsonData, err := json.Marshal(dataArr)
	if err != nil {
		return nil, err
	}

	return jsonData, err

}

func UploadFotoTempat(DataId int, FotoArr []model.FileD) ([]byte, error) {
	tx, err := Db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO foto_tempat (data_id, nama_foto) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for _, row := range FotoArr {
		_, err := stmt.Exec(DataId, row.Filename)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return []byte("Photos inserted"), nil

}

func GetTempatByID(id string) ([]byte, error) {
	type DataTempat struct {
		Id          int64   `json:"id"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Address     string  `json:"address"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
	}
	row := Db.QueryRow("SELECT * FROM data_tempat WHERE id = ?", id)

	var dataTempat DataTempat

	err := row.Scan(
		&dataTempat.Id,
		&dataTempat.Latitude,
		&dataTempat.Longitude,
		&dataTempat.Title,
		&dataTempat.Description,
		&dataTempat.UpdatedAt,
		&dataTempat.CreatedAt,
		&dataTempat.Address,
	)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(dataTempat)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
