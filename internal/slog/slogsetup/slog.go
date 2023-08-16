package slogsetup

import (
	"io"
	"log/slog"
	"strings"
	"sync"

	"gopkg.in/natefinch/lumberjack.v2"
)

var onceSetLogger sync.Once

func SetJSON(app, level string, w io.Writer) {
	handler := slog.NewJSONHandler(w,
		&slog.HandlerOptions{
			Level: parseLevel(level),
		},
	)

	logger := slog.New(handler)
	if app != "" {
		logger = logger.WithGroup(app)
	}

	onceSetLogger.Do(func() {
		slog.SetDefault(logger)
	})
}

func SetText(app, level string, w io.Writer) {
	handler := slog.NewTextHandler(w,
		&slog.HandlerOptions{
			Level: parseLevel(level),
		},
	)

	logger := slog.New(handler)
	if app != "" {
		logger = logger.WithGroup(app)
	}

	onceSetLogger.Do(func() {
		slog.SetDefault(logger)
	})
}

func WithRolling(filename string, w io.Writer) io.Writer {
	logger := &lumberjack.Logger{
		Filename:  filename,
		LocalTime: true,
	}
	return io.MultiWriter(w, logger)
}

func parseLevel(level string) slog.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
