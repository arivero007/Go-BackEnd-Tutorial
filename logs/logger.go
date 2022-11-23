package logs

import (
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func InitLogger() {
	file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error init log" + err.Error())
	}
	var flags = log.LstdFlags | log.Lshortfile
	InfoLogger = log.New(file, "INFO: ", flags)
	WarningLogger = log.New(file, "WARN: ", flags)
	ErrorLogger = log.New(file, "ERROR: ", flags)

	LogInfo("Custom logger initialized!")
}

func LogInfo(text string) {
	InfoLogger.Println(text)
}

func LogWarn(text string) {
	WarningLogger.Println(text)
}

func LogError(text string, err error) {
	ErrorLogger.Println(text + " " + err.Error())
}
