package logger

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/op/go-logging"
)

var (
	LogLevelStr  string
	logger       = logging.MustGetLogger("backhome")
	isLogFileSet = false
)

const LOG_FILE = ".backhome.log"

func InitLogger() {
	logfile, err := os.OpenFile(getLogFile(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to open log file: %v", err)
	} else {
		logging.SetBackend(logging.NewLogBackend(logfile, "", 0))
		isLogFileSet = true
	}

	acceptedLevels := map[string]logging.Level{
		"INFO":   logging.INFO,
		"DEBUG":  logging.DEBUG,
		"ERROR":  logging.ERROR,
		"SILENT": logging.CRITICAL + 1,
	}
	logLevel := logging.INFO
	if validLogLevel, exists := acceptedLevels[LogLevelStr]; exists {
		logLevel = validLogLevel
	}
	logging.SetLevel(logLevel, "backhome")

	logging.SetFormatter(logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05} %{color}%{level:.4s} %{id:03x}%{color:reset} %{message}`,
	))
}

func Info(msg string, args ...any) {
	log.Default()
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

func GetLogFile() string {
	if isLogFileSet {
		return getLogFile()
	}
	return "NOT SET!!!"
}

func getLogFile() string {
	home, _ := os.UserHomeDir()
	return strings.Join([]string{home, LOG_FILE}, string(os.PathSeparator))
}
