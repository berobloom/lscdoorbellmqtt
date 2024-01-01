package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetExecutableSourceDir() string {
	// get the directory where the executable is located /media/lscdoorbellmqtt
	dirPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Dir(dirPath)
}
