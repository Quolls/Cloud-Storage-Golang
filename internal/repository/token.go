package repository

import (
	// "database/sql"
	"fmt"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	db "github.com/Quolls/Cloud-Storage-Golang/internal/pkg/db/mysql"
)

func UpdateUserToken(user models.User, token string) bool {
	sqlStr := "REPLACE INTO token(`user_name`, `user_token`) VALUES(?, ?)"
	statement, err := db.GetDb().Prepare(sqlStr)
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer statement.Close()

	_, err = statement.Exec(user.Username, token)
	if err != nil {
		fmt.Println("Failed to execute statement, err:" + err.Error())
		return false
	}
	return true
}
