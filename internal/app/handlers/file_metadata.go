package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/services"
)

func GetFileMetadataHandler(c *gin.Context) {
	fileSha1 := c.Query("file_sha1")
	fileMetadata := services.GetFileMetadata(fileSha1)
	c.JSON(http.StatusOK, fileMetadata)
}
