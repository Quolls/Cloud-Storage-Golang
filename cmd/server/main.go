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
		v1.POST("/file", handlers.UploadHandler)
		v1.GET("/file/:name", handlers.DownloadHandler)
	}

	router.Run(":8080")
}
