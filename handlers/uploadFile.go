package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var FileExpirationMap = make(map[string]time.Time)

func UploadFile(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing form")
		return
	}

	file, handler, err := c.Request.FormFile("uploadFile")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading file")
		return
	}

	defer file.Close()

	if handler.Size > 10<<20 {
		c.String(http.StatusInternalServerError, "File size too large")
		return
	}

	fileid := uuid.New().String()
	filePath := fmt.Sprintf("./upload/%s", fileid)
	dst, err := os.Create(filePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error saving file")
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error writing file")
		return
	}

	downloadLink := fmt.Sprintf("/download/%s", fileid)
	c.String(http.StatusOK, "File uploaded successfully. Download at %s", downloadLink)

	FileExpirationMap[fileid] = time.Now().Add(24 * time.Hour)
}
