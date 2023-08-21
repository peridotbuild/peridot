package base

import (
	"fmt"
	"log"
	"strings"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
	LogLevelFatal LogLevel = "FATAL"
)

func Logf(level LogLevel, format string, args ...interface{}) {
	if !strings.HasPrefix(format, fmt.Sprintf("[%s] ", level)) {
		format = fmt.Sprintf("[%s]: %s", level, format)
	}

	if level == LogLevelFatal {
		log.Fatalf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func LogErrorf(format string, args ...interface{}) {
	Logf(LogLevelError, format, args...)
}

func LogWarnf(format string, args ...interface{}) {
	Logf(LogLevelWarn, format, args...)
}

func LogInfof(format string, args ...interface{}) {
	Logf(LogLevelInfo, format, args...)
}

func LogDebugf(format string, args ...interface{}) {
	Logf(LogLevelDebug, format, args...)
}

func LogFatalf(format string, args ...interface{}) {
	Logf(LogLevelFatal, format, args...)
}
