package logger

import (
	"log"
	"fmt"
	"os"
	"time"
)

type Level string

const (
	INFO     Level = "INFO"
	WARNING  Level = "WARNING"
	ERROR    Level = "ERROR"
	CRITICAL Level = "CRITICAL"
)

var (
	infoLogger     = log.New(os.Stdout, "", 0)
	warningLogger  = log.New(os.Stdout, "", 0)
	errorLogger    = log.New(os.Stdout, "", 0)
	criticalLogger = log.New(os.Stdout, "", 0)
)

func Log(level Level, message string, context ...interface{}) {
	timestamp := time.Now().Format("2025-04-14 11:00:00")

	var contextStr string
	if len(context) > 0 {
		contextStr = fmt.Sprintf(" | Context: %v", context)
	}

	logMessage := fmt.Sprintf("[%s] [%s] %s%s", timestamp, level, message, contextStr)

	switch level {
	case INFO:
		infoLogger.Println(logMessage)
	case WARNING:
		warningLogger.Println(logMessage)
	case ERROR:
		errorLogger.Println(logMessage)
	case CRITICAL:
		criticalLogger.Println(logMessage)
	default:
		infoLogger.Println(logMessage)
	}
}

func Info(message string, context ...interface{}) {
	Log(INFO, message, context...)
}

func Warning(message string, context ...interface{}) {
	Log(WARNING, message, context...)
}

func Error(message string, context ...interface{}) {
	Log(ERROR, message, context...)
}

func Critical(message string, context ...interface{}) {
	Log(CRITICAL, message, context...)
}
