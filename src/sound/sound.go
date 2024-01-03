package sound

import (
	"log"
	"lscdoorbellmqtt/utils"
	"os/exec"
)

func PlaySound(sound string) {
	cmd := exec.Command("/bin/sh", "/usr/sbin/playonspeaker.sh", "-f", getWaveFile(sound))
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getWaveFile(sound string) string {
	dirPath := utils.GetExecutableSourceDir()

	return dirPath + "/" + sound
}
