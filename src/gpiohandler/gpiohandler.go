package gpiohandler

import (
	"io/fs"
	"lscdoorbellmqtt/logger"
	"os"
	"strconv"
	"time"
)

const (
	BlueIndicator   = "50"
	RedIndicator    = "51"
	DoorbellBtnGPIO = "59"
	Max             = 3
	GPIOPath        = "/sys/class/gpio/gpio"
)

func SetHigh(pin string) {
	writeValue("/sys/class/gpio/gpio"+pin+"/value", "1")
}

func SetLow(pin string) {
	writeValue("/sys/class/gpio/gpio"+pin+"/value", "0")
}

func SetOutput(pin string) {
	writeDirection(pin, "out")
}

func SetInput(pin string) {
	writeDirection(pin, "in")
}

func ExportGPIO(pin string) {
	if _, err := os.Stat("/sys/class/gpio/gpio" + pin); os.IsNotExist(err) {
		writeValue("/sys/class/gpio/export", pin)
	}
}

func ReadGPIO(pin string) int {
	value, err := os.ReadFile("/sys/class/gpio/gpio" + pin + "/value")
	if err != nil {
		logger.Fatal(err.Error())
	}

	result, err := strconv.Atoi(string(value[:len(value)-1]))
	if err != nil {
		logger.Fatal(err.Error())
	}

	return result
}

func BellBlink() {
	for i := 0; i < Max; i++ {
		SetHigh(BlueIndicator)
		SetLow(RedIndicator)
		time.Sleep(time.Second)
		SetLow(BlueIndicator)
		SetHigh(RedIndicator)
		time.Sleep(time.Second)
	}
	SetLow(BlueIndicator)
	SetLow(RedIndicator)
}

func BootBlink() {
	SetHigh(BlueIndicator)
	SetHigh(RedIndicator)
	time.Sleep(time.Second)
	SetLow(RedIndicator)
	time.Sleep(time.Second)
	SetHigh(BlueIndicator)
	SetHigh(RedIndicator)
	time.Sleep(time.Second)
	SetLow(BlueIndicator)
	SetLow(RedIndicator)
}

func writeValue(path, value string) {
	err := os.WriteFile(path, []byte(value), fs.ModePerm)
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func writeDirection(pin string, direction string) {
	writeValue("/sys/class/gpio/gpio"+pin+"/direction", direction)
}

func GetBellState() int {
	return ReadGPIO(DoorbellBtnGPIO)
}

func Init() {
	ExportGPIO(BlueIndicator)
	ExportGPIO(RedIndicator)
	ExportGPIO(DoorbellBtnGPIO)

	SetOutput(BlueIndicator)
	SetOutput(RedIndicator)
	SetInput(DoorbellBtnGPIO)

	SetLow(BlueIndicator)
	SetLow(RedIndicator)
}
