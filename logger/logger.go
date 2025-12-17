package logger

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	warnLogger  = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime)
)

func Info(msg string) {
	infoLogger.Println(msg)
}

func Warn(msg string) {
	warnLogger.Println(msg)
}

func Error(msg string) {
	errorLogger.Println(msg)
}
