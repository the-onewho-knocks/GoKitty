package utils

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

var logLevel = new(slog.LevelVar)

func InitLogger(level string) {
	switch strings.ToLower(level) {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	case "error":
		logLevel.Set(slog.LevelError)
	default:
		logLevel.Set(slog.LevelInfo)
	}
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:       logLevel,
		TimeFormat:  time.RFC3339,
	})
	slog.SetDefault(slog.New(handler))
}

func LogDebug(msg string, args ...any) { slog.Debug(msg, args...) }
func LogInfo(msg string, args ...any)  { slog.Info(msg, args...) }
func LogWarn(msg string, args ...any)  { slog.Warn(msg, args...) }
func LogError(msg string, args ...any) { slog.Error(msg, args...) }
