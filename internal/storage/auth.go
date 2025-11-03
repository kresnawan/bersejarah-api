package storage

import (
	"app/internal/model"
)

func GetUserWithSpecificUsername(username string) ([]model.User, error) {
	var users []model.User
	rows, err := Db.Query("SELECT username, password from users WHERE username = ?", username)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}
