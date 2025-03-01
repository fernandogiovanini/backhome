package logger

import (
	"log"
	"log/slog"
	"os"
)

var LogLevelStr string

func InitLogger() {
	acceptedLevels := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
	}

	logLevel := slog.LevelInfo
	if validLogLevel, exists := acceptedLevels[LogLevelStr]; exists {
		logLevel = validLogLevel
	}

	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})),
	)
}

func Args(key string, value string) map[string]string {
	return map[string]string{key: value}
}

func transforArgs(args ...any) []any {
	var slogArgs []any
	for _, v := range args {
		if m, ok := v.(map[string]string); ok {
			for k, val := range m {
				slogArgs = append(slogArgs, slog.String(k, val))
			}
			continue
		}
		slogArgs = append(slogArgs, v)
	}
	return slogArgs
}

func Info(msg string, args ...any) {
	slog.Info(msg, transforArgs(args...)...)
}

func Err(err error, args ...any) {
	slog.Error(err.Error(), transforArgs(args...)...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, transforArgs(args...)...)
}

func Fatalf(msg string, args ...any) {
	log.Fatalf(msg, args...)
}
