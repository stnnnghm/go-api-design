package logger

import (
	"log"
	"os"
)

var (
	Info  = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	Error = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)
