package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

func CalculateFileSha1(file *multipart.FileHeader) (string, error) {

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, src); err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	sha1 := hex.EncodeToString(hashBytes)

	return sha1, nil
}

func EncodeString(input string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Error comparing passwords: ", err)
		return false
	}
	return true
}
