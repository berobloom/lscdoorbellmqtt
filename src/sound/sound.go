package sound

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func PlaySound() {
	cmd := exec.Command("/bin/sh", "/usr/sbin/playonspeaker.sh", "-f", getWaveFile())
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getWaveFile() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("Error:", err)
	}

	dirPath := filepath.Dir(exePath)

	return dirPath + "/dingdong.wav"
}
