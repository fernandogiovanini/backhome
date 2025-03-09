package logger

import (
	"github.com/op/go-logging"
)

var LogLevelStr string
var logger = logging.MustGetLogger("backhome")

func InitLogger() {
	acceptedLevels := map[string]logging.Level{
		"INFO":  logging.Level(logging.INFO),
		"DEBUG": logging.Level(logging.DEBUG),
		"ERROR": logging.Level(logging.ERROR),
	}
	logLevel := logging.INFO
	if validLogLevel, exists := acceptedLevels[LogLevelStr]; exists {
		logLevel = validLogLevel
	}

	logging.SetFormatter(logging.MustStringFormatter(
		`%{color}%{level:.4s} %{id:03x}%{color:reset} %{message}`,
	))
	logging.SetLevel(logLevel, "backhome")
}

func Info(msg string, args ...any) {
	logger.Infof(msg, args...)
}

func Err(msg string, args ...any) {
	logger.Errorf(msg, args...)
}

func Debug(msg string, args ...any) {
	logger.Debugf(msg, args...)
}

func Fatalf(msg string, args ...any) {
	logger.Fatalf(msg, args...)
}
