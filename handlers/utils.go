package handlers

import (
	"fmt"
	"time"
)

// Returns the path of the file on the server
func GetFilePathFromID(fileID string) string {
	return fmt.Sprintf("./uploads/%s", fileID)
}

// Returns the expiration time for a file
func GetFileExpirationTime(fileID string) time.Time {
	return FileExpirationMap[fileID]
}
