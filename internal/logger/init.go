package logger

import (
	"log/slog"
	"os"
)

func InitLogger() *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	return slog.New(handler)
}
