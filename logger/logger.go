package logger

import (
	"strings"

	"github.com/op/go-logging"
)

var LogLevelStr string
var localLogger = logging.MustGetLogger("backhome")

func InitLogger() {
	acceptedLevels := map[string]logging.Level{
		"INFO":  logging.Level(logging.INFO),
		"DEBUG": logging.Level(logging.DEBUG),
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

// An Attr is a key-value pair.
type Argument struct {
	Key   string
	Value string
}

func Args(key string, value string) Argument {
	return Argument{Key: key, Value: value}
}

func transforArgs(args ...any) []any {
	var returnArgs []any
	for _, v := range args {
		if a, ok := v.(Argument); ok {
			returnArgs = append(returnArgs, strings.Join([]string{a.Key, a.Value}, "="))
			continue
		}
		returnArgs = append(returnArgs, v)
	}
	return returnArgs
}

func addPlaceholdersToMessage(msg string, args ...any) string {
	finalMessage := []string{msg}
	num := len(args) - strings.Count(msg, "%s")
	for i := 0; i < num; i++ {
		finalMessage = append(finalMessage, "%s")
	}
	return strings.Join(finalMessage, " ")
}

func Info(msg string, args ...any) {
	localLogger.Infof(addPlaceholdersToMessage(msg, args...), transforArgs(args...)...)
}

func Err(err error, args ...any) {
	localLogger.Errorf(addPlaceholdersToMessage(err.Error(), args...), transforArgs(args...)...)
}

func Debug(msg string, args ...any) {
	localLogger.Debugf(addPlaceholdersToMessage(msg, args...), transforArgs(args...)...)
}

func Fatalf(msg string, args ...any) {
	localLogger.Fatalf(addPlaceholdersToMessage(msg, args...), transforArgs(args...)...)
}
