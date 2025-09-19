package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func PathGenerator(pathName string) string {
	ext := ""

	tempDir := os.TempDir()
	pathName = "/" + pathName

	if runtime.GOOS == "windows" {
		pathName = strings.Replace(pathName, "/", "\\", -1)
		ext = ".exe"
	}

	return fmt.Sprintf("%s%s%s", tempDir, pathName, ext)
}
