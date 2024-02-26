package repository

import (
	"database/sql"
	"fmt"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	db "github.com/Quolls/Cloud-Storage-Golang/internal/pkg/db/mysql"
)

func InsertUser(username, password string) bool {
	sqlStr := "INSERT IGNORE INTO user(user_name, user_pwd) values(?, ?)"

	statement, err := db.GetDb().Prepare(sqlStr)
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer statement.Close()

	result, err := statement.Exec(username, password)
	if err != nil {
		fmt.Println("Failed to execute statement, err:" + err.Error())
		return false
	}
	if rf, err := result.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("User with username:%s has been signed up before\n", username)
			return false
		}
		return true
	}
	return false
}

func GetUser(username, password string) (*models.User, error) {
	sqlStr := "SELECT user_name, user_pwd, email, phone, email_validated, phone_validated, " +
		" signup_at, last_active, profile, status " +
		" FROM user WHERE user_name = ? LIMIT 1"

	statement, err := db.GetDb().Prepare(sqlStr)
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return &models.User{}, err
	}
	defer statement.Close()

	user := models.User{}
	err = statement.QueryRow(username).Scan(&user.Username, &user.Password, &user.Email, &user.Phone,
		&user.EmailValidated, &user.PhoneValidated, &user.SignUpAt, &user.LastActive, &user.Profile, &user.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("User with username: %s is not found\n", username)
			return &models.User{}, nil
		} else {
			fmt.Printf("Failed to execute statement, err: %s\n", err.Error())
			return &models.User{}, err
		}
	}

	return &user, nil
}

// 	if rows == nil {
// 		fmt.Printf("User with username:%s is not found\n", username)
// 		return false
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		var pwd string
// 		err = rows.Scan(&username, &pwd)
// 		if err != nil {
// 			fmt.Println("Failed to scan row, err:" + err.Error())
// 			return false
// 		}
// 		if pwd != password {
// 			fmt.Printf("Password for user with username:%s is not correct\n", username)
// 			return false
// 		}
// 	}

// 	return false
