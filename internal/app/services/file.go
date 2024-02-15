package services

import (
	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
)

var fileMetadataCollections map[string]models.FileMetadata

func init() {
	fileMetadataCollections = make(map[string]models.FileMetadata)
}

func UpdateFileMetadata(metadata models.FileMetadata) {
	fileMetadataCollections[metadata.FileSha1] = metadata
}

func GetFileMetadata(fileSha1 string) models.FileMetadata {
	return fileMetadataCollections[fileSha1]
}
