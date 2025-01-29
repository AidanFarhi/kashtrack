package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger(logFile string) {
	var output *os.File
	if logFile == "local" {
		output = os.Stdout
	} else {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		output = file
	}
	Logger = log.New(output, "", log.Ldate|log.Ltime|log.Lshortfile)
}
