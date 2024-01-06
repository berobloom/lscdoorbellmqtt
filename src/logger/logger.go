package logger

import (
	"io"
	"log"
	"lscdoorbellmqtt/utils"
	"os"
)

type Level int

const (
	INFO Level = iota
	ERROR
)

var (
	Info   *log.Logger
	Error  *log.Logger
	Status *log.Logger
	level  Level
)

func Init(logLevel Level) {
	dirPath := utils.GetExecutableSourceDir()
	logFile, err := os.OpenFile(dirPath+"lscdoorbellmqtt.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	level = logLevel

	Info = log.New(io.MultiWriter(os.Stdout, logFile), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stdout, logFile), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Status = log.New(io.MultiWriter(os.Stdout, logFile), "STATUS: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Infof(format string, v ...interface{}) {
	if level <= INFO {
		Info.Printf(format, v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if level <= ERROR {
		Error.Printf(format, v...)
	}
}

func Fatal(format string, v ...interface{}) {
	Error.Printf(format, v...)
	os.Exit(1)
}
