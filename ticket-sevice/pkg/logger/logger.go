package logger

import (
	"log"
	"path/filepath"
	"runtime"
)

func LogStatus(status string, message string) {
	// Get caller info: 1 = the function that called LogStatus
	_, fullPath, line, ok := runtime.Caller(1)
	if !ok {
		log.Println("Could not get caller information")
		return
	}

	// Extract file name and folder name
	fileName := filepath.Base(fullPath)
	dirName := filepath.Base(filepath.Dir(fullPath))

	log.Printf("[%s] %s/%s:%d - %s\n", status, dirName, fileName, line, message)
}
