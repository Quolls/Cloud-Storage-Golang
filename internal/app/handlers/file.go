package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/models"
	"github.com/Quolls/Cloud-Storage-Golang/internal/app/util"
)

func UploadHandler(c *gin.Context) {
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
	models.UpdateFileMetadata(fileMetadata)

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully!", "filename": file.Filename})
}

func DownloadHandler(c *gin.Context) {

	name := c.Param("name")
	fmt.Println("Received file: ", name)

	if _, err := os.Stat("./tmp/" + name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found!"})
		return
	}

	c.FileAttachment("./tmp/"+name, name)
}
