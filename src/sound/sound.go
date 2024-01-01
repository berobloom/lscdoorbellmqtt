package sound

import (
	"log"
	"lscdoorbellmqtt/utils"
	"os/exec"
)

func PlaySound() {
	cmd := exec.Command("/bin/sh", "/usr/sbin/playonspeaker.sh", "-f", getWaveFile())
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getWaveFile() string {
	dirPath := utils.GetExecutableSourceDir()

	return dirPath + "/dingdong.wav"
}
