package loggy

import (
	"log"
	"os"
)

var info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// Info logs information events
func Info(v ...interface{}) {
	info.Println(v)
}

var warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

// Warn logs warning events
func Warn(v ...interface{}) {
	warn.Println(v)
}

var error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

// Error logs error events
func Error(v ...interface{}) {
	error.Println(v)
}

// Fatal same as Error call followed by os.Exit(1)
func Fatal(v ...interface{}) {
	Error(v)
	os.Exit(1)
}
