package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/services"
)

func GetFileMetadataHandler(c *gin.Context) {
	fileSha1 := c.Query("file_sha1")

	if fileSha1 == "" {
		fileMetadataCollections := services.GetFileMetadataByRange("all")
		c.JSON(http.StatusOK, fileMetadataCollections)
		return
	}

	fileMetadata := services.GetFileMetadata(fileSha1)
	if fileMetadata.FileSha1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
		return
	}

	c.JSON(http.StatusOK, fileMetadata)
}

func UpdateFileMetadataHandler(c *gin.Context) {

	fileSha1 := c.Query("file_sha1")
	newFileName := c.Query("file_name")

	curFilemetadata := services.GetFileMetadata(fileSha1)
	fmt.Print(fileSha1)
	fmt.Print(curFilemetadata)
	if curFilemetadata.FileSha1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
		return
	}
	curFilemetadata.FileName = newFileName

	services.UpdateFileMetadata(curFilemetadata)
	c.JSON(http.StatusOK, curFilemetadata)
}

func DeleteFileMetadataHandler(c *gin.Context) {
	fileSha1 := c.Params.ByName("file_sha1")

	curFilemetadata := services.GetFileMetadata(fileSha1)
	if curFilemetadata.FileSha1 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
		return
	}

	services.DeleteFileMetadata(fileSha1)
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully!"})
}
