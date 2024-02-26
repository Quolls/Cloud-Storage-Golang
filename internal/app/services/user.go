package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

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

	token, err := GenerateToken(user)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return false
	}
	if !repository.UpdateUserToken(user, token) {
		fmt.Println("Failed to update user token!")
		return false
	}

	return true
}

func GenerateToken(user models.User) (string, error) {
	expireationTime := time.Now().Add(1 * time.Hour)
	fmt.Println(expireationTime)

	claims := &jwt.StandardClaims{
		ExpiresAt: expireationTime.Unix(),
		Issuer:    "Quolls",
		Subject:   user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}
