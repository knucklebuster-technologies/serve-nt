package loggy

import (
	"io"
	"log"
	"os"
)

// Info logs information events
var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime).Println

// Warn logs warning events
var Warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime).Println

// Error logs error events
var Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime).Println

// Fatal same as Error call followed by os.Exit(1)
var Fatal = log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime).Fatalln

// Set changes the default io.Writer with one supplied by the caller
func Set(w io.Writer) {
	Info = getLogger(w, "INFO: ")
	Warn = getLogger(w, "WARN: ")
	Error = getLogger(w, "ERROR: ")
	Fatal = getLogger(w, "FATAL: ")
}

func getLogger(w io.Writer, prefix string) func(...interface{}) {
	if prefix == "FATAL: " {
		return log.New(w, prefix, log.Ldate|log.Ltime).Fatalln
	}
	return log.New(w, prefix, log.Ldate|log.Ltime).Println
}