package main

import (
	"lscdoorbellmqtt/config"
	"lscdoorbellmqtt/gpiohandler"
	"lscdoorbellmqtt/logger"
	"lscdoorbellmqtt/mqtt"
	"lscdoorbellmqtt/sound"
)

func main() {
	config.Init()

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

	gpiohandler.Init()
	mqtt.Start()
}
