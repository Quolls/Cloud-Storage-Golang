package handlers

import (
	"fmt"
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
	fmt.Println(econdedPassword)
	fmt.Println(util.ComparePasswords(econdedPassword, password))
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

func SignInUserHandler(c *gin.Context) {
	user := models.User{
		Username: c.PostForm("user_name"),
		Password: c.PostForm("user_pwd"),
	}

	if len(user.Username) < 3 || len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username must be at least 3 characters and password must be at least 6 characters!"})
		return
	}

	if !services.SignInUser(user) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign in user!"})
		return
	}

	// TODO: Generate JWT token and send it back to the client

	c.JSON(http.StatusOK, gin.H{"message": "User signed in successfully!"})
}
