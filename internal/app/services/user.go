package services

import (
	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	"github.com/Quolls/Cloud-Storage-Golang/internal/repository"
)

func SignUpUser(user models.User) bool {
	return repository.InsertUser(user.Username, user.Password)
}
