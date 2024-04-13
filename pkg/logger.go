package pkg

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

type Logger struct {
	level LogLevel
}

func NewLogger(level LogLevel) *Logger {
	return &Logger{level: level}
}

func (l *Logger) log(level LogLevel, message string, args ...interface{}) {
	if l.level <= level {
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf(message, args...)
		fmt.Printf("[%s] %s: %s\n", timeNow, levelString(level), formattedMessage)
	}
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log(INFO, message, args...)
}

func (l *Logger) Warning(message string, args ...interface{}) {
	l.log(WARNING, message, args...)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.log(ERROR, message, args...)
}

func levelString(level LogLevel) string {
	switch level {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
