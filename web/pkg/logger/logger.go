package logger

import (
	"log/slog"
	"os"
	"strings"

	"github.com/Pineapple217/Sortify/web/pkg/config"
)

func NewLogger(set config.Logger) *slog.Logger {
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: GetLogLevel(set.Level),
	})
	l := slog.New(h)
	return l
}

func GetLogLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
