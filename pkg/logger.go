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

func (l *Logger) log(level LogLevel, message string) {
	if l.level <= level {
		fmt.Printf("[%s] %s: %s\n", time.Now().Format("2006-01-02 15:04:05"), levelString(level), message)
	}
}

func (l *Logger) Info(message string) {
	l.log(INFO, message)
}

func (l *Logger) Warning(message string) {
	l.log(WARNING, message)
}

func (l *Logger) Error(message string) {
	l.log(ERROR, message)
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
