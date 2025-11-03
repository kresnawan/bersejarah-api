package storage

import (
	"app/internal/model"
	"encoding/json"
	"log"
)

func GetAllUsers() []byte {
	var users []model.User
	rows, err := Db.Query("SELECT id, name, username FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Username); err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

func RegisterUser(name string, username string, password string) (int64, error) {
	var query string = "INSERT INTO `users` (`username`, `password`, `name`) VALUES (?, ?, ?)"
	insertResult, err := Db.Exec(query, username, password, name)
	if err != nil {
		return 0, err
	}

	insertId, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertId, nil
}
