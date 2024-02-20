package repository

import (
	"fmt"

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
