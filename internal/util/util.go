package util

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"mime/multipart"
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
