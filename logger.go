package main

import (
	"log"
	"os"
)

var linfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var lwarn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
var lerror = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
