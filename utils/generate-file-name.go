package utils

import (
	"fmt"
	"time"
)

func GenerateFileName(extension string) string {
	now := time.Now()
	fileName := fmt.Sprintf("%02d_%02d_%02d_%02d_%02d_%d%s",
		now.Second(), now.Minute(), now.Hour(), now.Day(), now.Month(), now.Year(), extension)
	return fileName
}
