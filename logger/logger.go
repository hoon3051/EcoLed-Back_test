package logger

import (
	"log"
	"os"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func init() {
	// TODO : Logging msg to File
	Debug = log.New(os.Stdout, "[DEBUG  ] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "[INFO   ] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR  ] ", log.Ldate|log.Ltime|log.Lshortfile)
}
