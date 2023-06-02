package logger

import (
	"log"
	"os"
)

type Logger interface {
	Error(msg ...interface{})
	Info(msg ...interface{})
}

type Log struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

func NewLog() *Log {
	return &Log{
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
	}
}

func (l *Log) Error(msg ...interface{}) {
	l.errorLogger.Println(msg...)
}

func (l *Log) Info(msg ...interface{}) {
	l.infoLogger.Println(msg...)
}
