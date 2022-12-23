package dao

import (
	"fmt"

	"Q-A/model"
)

func UpdatePassword(username, newPassword string) error {
	_, err := dB.Exec("UPDATE user SET password = ? WHERE username = ?", newPassword, username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("服务器出错")
			return
		}
	}()

	row := dB.QueryRow("SELECT id,password FROM user WHERE username = ?", username)
	if row.Err() != nil {
		return user, row.Err()
	}

	if err := row.Scan(&user.Id, &user.Password); err != nil {
		return user, err
	}
	return user, nil
}

func InsertUser(user model.User) error {
	_, err := dB.Exec("INSERT INTO user(username,password) values(?,?)", user.Username, user.Password)
	return err
}

func DelUser(username string) error {
	_, err := dB.Exec("DELETE FROM user WHERE username = ?", username)
	return err
}
