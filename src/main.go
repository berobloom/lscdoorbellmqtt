package main

import (
	"lscdoorbellmqtt/config"
	"lscdoorbellmqtt/gpiohandler"
	"lscdoorbellmqtt/logger"
	"lscdoorbellmqtt/mqtt"
	"lscdoorbellmqtt/openipc"
	"lscdoorbellmqtt/sound"
)

func main() {
	openipc.GetOSReleaseVariables()
	config.Init()
	gpiohandler.Init()

	logLevel := config.GetString("settings.log_level")
	switch logLevel {
	case "INFO":
		logger.Init(logger.INFO)
	case "ERROR":
		logger.Init(logger.ERROR)
	default:
		logger.Init(logger.INFO)
	}

	go sound.PlaySound("boot.wav")
	go gpiohandler.BootBlink()

	logger.Status.Println("Starting lscdoorbellmqtt...")

	mqtt.Start()
}
