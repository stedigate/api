package logger

import (
	"log/slog"
)

func NewConfig() *Config {
	return &Config{}
}

func NewLogger(h slog.Handler) *slog.Logger {
	logger := slog.New(h)
	slog.SetDefault(logger)
	return logger
}
