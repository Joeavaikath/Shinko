package logger

import (
	"log"
	"os"
	"sync"
)

// LogLevel defines the level of logging
type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

// Logger represents a logger with different log levels
type Logger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

var (
	instance *Logger
	once     sync.Once
)

// initLogger initializes the singleton instance of the Logger
func initLogger() {
	instance = &Logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// getInstance returns the singleton instance of the Logger
func getInstance() *Logger {
	once.Do(initLogger)
	return instance
}

// Info logs an info message
func Info(v ...interface{}) {
	getInstance().infoLogger.Println(v...)
}

// Warning logs a warning message
func Warning(v ...interface{}) {
	getInstance().warningLogger.Println(v...)
}

// Error logs an error message
func Error(v ...interface{}) {
	getInstance().errorLogger.Println(v...)
}
