package openipc

import (
	"log"

	"github.com/joho/godotenv"
)

func GetOSReleaseVariables() {
	err := godotenv.Load("/etc/os-release")
	if err != nil {
		log.Fatal("Error loading /etc/os-release file")
	}
}
