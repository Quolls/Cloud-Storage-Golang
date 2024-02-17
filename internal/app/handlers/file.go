package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	"github.com/Quolls/Cloud-Storage-Golang/internal/app/services"
	"github.com/Quolls/Cloud-Storage-Golang/internal/util"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received file: ", file.Filename)
	path := "./tmp/" + file.Filename
	fmt.Println("Saving file to:", path)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileSha1, err := util.CalculateFileSha1(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileMetadata := models.FileMetadata{
		FileSha1: fileSha1,
		FileName: file.Filename,
		FileSize: file.Size,
		FilePath: path,
		CreateAt: time.Now().String(),
	}
	services.UpdateFileMetadata(fileMetadata)
	fmt.Println("File metadata:", fileMetadata)
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!", "filename": file.Filename})
}

func DownloadFileHandler(c *gin.Context) {

	fileSha1 := c.Query("file_sha1")
	fileMetadata := services.GetFileMetadata(fileSha1)

	if _, err := os.Stat(fileMetadata.FilePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
		return
	}

	c.FileAttachment(fileMetadata.FilePath, fileMetadata.FileName)
}

func DeleteFileHandler(c *gin.Context) {
	fileSha1 := c.Query("file_sha1")
	fileMetadata := services.GetFileMetadata(fileSha1)

	if _, err := os.Stat(fileMetadata.FilePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
		return
	}

	err := os.Remove(fileMetadata.FilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	services.DeleteFileMetadata(fileSha1)
	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully!"})
}
