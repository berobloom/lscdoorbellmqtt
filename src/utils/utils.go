package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetExecutableSourceDir() string {
	dirPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Dir(dirPath)
}
