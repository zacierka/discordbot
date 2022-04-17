package logger

import (
	"log"
	"os"
)

var ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
var InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
