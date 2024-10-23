package handlers

import (
	"fmt"
	"log"
	"os"
	"time"
)

func CleanupExpiredFiles() {
	for {
		time.Sleep(1 * time.Hour)
		for fileid, expirationTime := range FileExpirationMap {
			if time.Now().After(expirationTime) {
				deleteFile(fileid)
			}
		}
	}

}

func deleteFile(fileID string) {
	filePath := fmt.Sprintf("./upload/%s", fileID)
	if err := os.Remove(filePath); err != nil {
		log.Printf("failed to delete the file %s: %v", fileID, err)
	} else {
		delete(FileExpirationMap, fileID)
		log.Printf("Deleted expired file %s", fileID)
	}

}
