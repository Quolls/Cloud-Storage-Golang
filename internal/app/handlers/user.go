package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	"github.com/Quolls/Cloud-Storage-Golang/internal/app/services"
	"github.com/Quolls/Cloud-Storage-Golang/internal/util"
)

func SignUpUserHandler(c *gin.Context) {

	username := c.PostForm("user_name")
	password := c.PostForm("user_pwd")

	if len(username) < 3 || len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username must be at least 3 characters and password must be at least 6 characters!"})
		return
	}

	econdedPassword, err := util.EncodeString(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode password!"})
	}

	user := models.User{
		Username: username,
		Password: econdedPassword,
	}

	if !services.SignUpUser(user) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up user!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully!"})
}
