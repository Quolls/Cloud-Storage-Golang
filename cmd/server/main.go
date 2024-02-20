package main

import (
	// "fmt"
	// "net/http"
	// "log"

	"github.com/gin-gonic/gin"

	"github.com/Quolls/Cloud-Storage-Golang/internal/app/handlers"
)

func main() {

	// cfg, err := cfg.Load()

	// if err != nil {
	// 	log.Fatal("Error loading cofig: %v", err)
	// }

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/file", handlers.UploadFileHandler)
		v1.GET("/file", handlers.DownloadFileHandler)
		v1.DELETE("/file", handlers.DeleteFileHandler)

		v1.GET("/file/metadata", handlers.GetFileMetadataHandler)
		v1.PUT("/file/metadata", handlers.UpdateFileMetadataHandler)

		v1.POST("/user", handlers.SignUpUserHandler)
	}

	router.Run(":8080")
}
