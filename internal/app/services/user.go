package services

import (
	"fmt"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	"github.com/Quolls/Cloud-Storage-Golang/internal/repository"
	"github.com/Quolls/Cloud-Storage-Golang/internal/util"
)

func SignUpUser(user models.User) bool {
	return repository.InsertUser(user.Username, user.Password)
}

func SignInUser(user models.User) bool {

	userData, _ := repository.GetUser(user.Username, user.Password)
	if !util.ComparePasswords(userData.Password, user.Password) {
		fmt.Println("Password is not correct!")
		return false
	}

	return true
}
