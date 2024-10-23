package main

import (
	"fileHandler/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// File upload route
	r.POST("/upload", handlers.UploadFile)

	// File download route
	r.GET("/download/:id", handlers.DownloadFile)

	// Start the cleanup Goroutine for expired files
	go handlers.CleanupExpiredFiles()

	// Start the server
	log.Println("Server started on :8080")
	r.Run(":8080")
}
