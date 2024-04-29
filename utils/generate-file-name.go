package utils

import (
	"fmt"
	"time"
)

func GenerateFileName(extension string) string {
	now := time.Now()
	// Format the time to include nanoseconds
	fileName := fmt.Sprintf("%09d_%02d_%02d_%02d_%02d_%02d_%d_%s",
		now.Nanosecond(), now.Second(), now.Minute(), now.Hour(), now.Day(), int(now.Month()), now.Year(), extension)
	return fileName
}
