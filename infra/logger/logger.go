package logger

import (
	"log"
	"os"
)

// ANSI escape codes for colors
const (
	red    = "\033[31m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

var (
	stdoutError *log.Logger
	fileError   *log.Logger
	stdoutInfo  *log.Logger
	fileInfo    *log.Logger
	stdoutWarn  *log.Logger
	fileWarn    *log.Logger
)

func Init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fileError = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	stdoutError = log.New(os.Stdout, red+"ERROR: "+reset, 0)

	fileInfo = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	stdoutInfo = log.New(os.Stdout, blue+"INFO: "+reset, 0)

	fileWarn = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	stdoutWarn = log.New(os.Stdout, yellow+"WARN: "+reset, 0)
	stdoutInfo.Println("Logger initialized")
}

func Error(v ...any) {
	fileError.Println(v...)
	stdoutError.Fatal(v...)
}

func Info(v ...any) {
	fileInfo.Println(v...)
	stdoutInfo.Println(v...)
}

func Warn(v ...any) {
	fileWarn.Println(v...)
	stdoutWarn.Println(v...)
}
