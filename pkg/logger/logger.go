package logger

import (
	"log"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

func Success(msg string) {
	log.Println(colorGreen + "[SUCCESS] " + msg + colorReset)
}

func Error(msg string) {
	log.Println(colorRed + "[ERROR] " + msg + colorReset)
}

func Warn(msg string) {
	log.Println(colorYellow + "[WARN] " + msg + colorReset)
}

func Info(msg string) {
	log.Println(colorBlue + "[INFO] " + msg + colorReset)
}

func Fatal(msg string, err error) {
	log.Fatalf(colorRed+"[FATAL] %s: %v"+colorReset, msg, err)
}
