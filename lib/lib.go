package lib

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type LogOptions struct {
	Level  string // "debug" | "info" | "warn" | "error"
	Format string // "json" | "text"
}

func Init(opt LogOptions) {
	level := parseLevel(opt.Level)

	var h slog.Handler
	if opt.Format == "text" {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	} else {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}

	slog.SetDefault(slog.New(h))
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "":
		return slog.LevelInfo
	default:
		return slog.LevelInfo
	}
}

func MustParseLevel(level string) (slog.Level, error) {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("unknown log level %q", level)
	}
}
