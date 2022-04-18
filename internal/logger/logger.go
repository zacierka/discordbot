package logger

import (
	"log"
	"os"
	"strings"
)

var ErrorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
var InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)

var TestLog = log.New(os.Stdout, "[TEST]\t", log.Ldate|log.Ltime|log.Lshortfile)

var sb strings.Builder

func LOGMSG(str ...string) {
	for _, val := range str {
		sb.WriteString(val)
	}
	InfoLog.Println(sb.String())
	sb.Reset()
}

func LOGERR(str ...string) {
	for _, val := range str {
		sb.WriteString(val)
	}
	ErrorLog.Println(sb.String())
	sb.Reset()
}

func LOGFATAL(str ...string) {
	for _, val := range str {
		sb.WriteString(val)
	}
	ErrorLog.Fatalln(sb.String())
}
