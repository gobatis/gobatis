package logger

import "log"

func Debug(v ...interface{}) {
	log.Println(v...)
}
