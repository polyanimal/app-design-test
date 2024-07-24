package logger

import (
	"fmt"
	"log"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger(logger *log.Logger) Logger {
	return Logger{
		logger: logger,
	}
}

func (l *Logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.logger.Printf("[Error]: %s\n", msg)
}

func (l *Logger) LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.logger.Printf("[Info]: %s\n", msg)
}
