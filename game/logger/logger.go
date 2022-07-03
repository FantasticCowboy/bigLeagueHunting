package logger

import (
	"log"
)

var std = log.Default()

func Logf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
