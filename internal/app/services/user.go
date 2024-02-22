package services

import (
	"fmt"
	"time"

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

	token := GenerateToken(user)
	if !repository.UpdateUserToken(user, token) {
		fmt.Println("Failed to update user token!")
		return false
	}

	return true
}

func GenerateToken(user models.User) string {
	fmt.Println(time.Now().Unix())

	return "token:123456"
}
