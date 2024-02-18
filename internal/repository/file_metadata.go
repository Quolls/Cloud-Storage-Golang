package repository

import (
	"fmt"

	db "github.com/Quolls/Cloud-Storage-Golang/internal/pkg/db/mysql"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
)

func InsertFileMetadata(fileMetadata models.FileMetadata) bool {
	sqlStr := "INSERT IGNORE INTO file_metadata(file_sha1, file_name, file_size, file_path, status) values(?,?,?,?,1)"
	statement, err := db.GetDb().Prepare(sqlStr)

	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer statement.Close()

	result, err := statement.Exec(fileMetadata.FileSha1, fileMetadata.FileName, fileMetadata.FileSize, fileMetadata.FilePath)
	if err != nil {
		fmt.Println("Failed to execute statement, err:" + err.Error())
		return false
	}
	if rf, err := result.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded before\n", fileMetadata.FileSha1)
		}
		return true
	}
	return false
}

func GetFileMetadata(fileSha1 string) (*models.FileMetadata, error) {
	sqlStr := "SELECT file_sha1, file_name, file_size, file_path, status " +
		"FROM file_metadata WHERE file_sha1 = ? AND status = 1"

	statement, err := db.GetDb().Prepare(sqlStr)
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return &models.FileMetadata{}, err
	}
	defer statement.Close()

	fileMetadata := models.FileMetadata{}
	err = statement.QueryRow(fileSha1).Scan(&fileMetadata.FileSha1, &fileMetadata.FileName,
		&fileMetadata.FileSize, &fileMetadata.FilePath, &fileMetadata.Status)
	if err != nil {
		fmt.Println("Failed to execute statement, err:" + err.Error())
		return &models.FileMetadata{}, err
	}
	return &fileMetadata, nil
}
