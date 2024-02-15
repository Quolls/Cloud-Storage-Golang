package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
