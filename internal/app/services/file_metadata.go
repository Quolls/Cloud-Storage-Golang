package services

import (
	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	"github.com/Quolls/Cloud-Storage-Golang/internal/repository"
)

// var fileMetadataCollections map[string]models.FileMetadata

func init() {
	// fileMetadataCollections = make(map[string]models.FileMetadata)
}

func UpdateFileMetadata(metadata models.FileMetadata) bool {
	return repository.UpdateFileMetadata(metadata)
}

func InsertFileMetadataToDB(metadata models.FileMetadata) bool {
	return repository.InsertFileMetadata(metadata)
}

// func GetFileMetadata(fileSha1 string) models.FileMetadata {
// 	return fileMetadataCollections[fileSha1]
// }

func GetFileMetadataFromDB(fileSha1 string) (models.FileMetadata, error) {
	filemetadata, err := repository.GetFileMetadata(fileSha1)
	if err != nil {
		return models.FileMetadata{}, err
	}
	return *filemetadata, nil
}

// func GetFileMetadataByRange(timeRange string) map[string]models.FileMetadata {
// 	return fileMetadataCollections
// }

func DeleteFileMetadataFromDB(fileSha1 string) bool {
	return repository.DeleteFileMetadata(fileSha1)
}
