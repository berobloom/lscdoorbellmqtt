package main

import (
	"lscdoorbellmqtt/config"
	"lscdoorbellmqtt/gpiohandler"
	"lscdoorbellmqtt/mqtt"
)

func main() {
	config.Init()
	gpiohandler.Init()
	mqtt.Start()
}
