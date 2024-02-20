package handlers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	"github.com/Quolls/Cloud-Storage-Golang/internal/app/services"
)

func GetFileMetadataHandler(c *gin.Context) {
	fileSha1 := c.Query("file_sha1")

	if fileSha1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file_sha1 is required!"})
		return
	}

	fileMetadata, err := services.GetFileMetadataFromDB(fileSha1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fileMetadata)
}

func UpdateFileMetadataHandler(c *gin.Context) {

	fileSha1 := c.Query("file_sha1")
	newFileName := c.Query("file_name")

	curFilemetadata, err := services.GetFileMetadataFromDB(fileSha1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if curFilemetadata == (models.FileMetadata{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
		return
	}
	curFilemetadata.FileName = newFileName

	if !services.UpdateFileMetadata(curFilemetadata) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update file metadata!"})
		return
	}
	c.JSON(http.StatusOK, curFilemetadata)
}

// func DeleteFileMetadataHandler(c *gin.Context) {
// 	fileSha1 := c.Query("file_sha1")

// 	curFilemetadata := services.GetFileMetadata(fileSha1)
// 	if curFilemetadata.FileSha1 == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
// 		return
// 	}

// 	services.DeleteFileMetadataFromDB(fileSha1)
// 	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully!"})
// }
