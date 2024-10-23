package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	fileId := c.Param("id")
	filePath := GetFilePathFromID(fileId)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	c.File(filePath)
}
