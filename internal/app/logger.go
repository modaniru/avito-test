package app

import (
	log "log/slog"
	"os"
)

func configureLogger(level string) {
	lvl := log.LevelDebug
	switch level {
	case "info":
		lvl = log.LevelInfo
	case "warn":
		lvl = log.LevelWarn
	case "error":
		lvl = log.LevelError
	}
	logger := log.New(log.NewJSONHandler(os.Stdout, &log.HandlerOptions{Level: lvl}))
	log.SetDefault(logger)
}
