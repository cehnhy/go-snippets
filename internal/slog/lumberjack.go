package slog

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

func WithRolling(filename string, w io.Writer) io.Writer {
	logger := &lumberjack.Logger{
		Filename:  filename,
		LocalTime: true,
	}
	return io.MultiWriter(w, logger)
}
