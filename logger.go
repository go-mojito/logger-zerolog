package zerolog

import (
	"log/slog"

	slogzerolog "github.com/samber/slog-zerolog/v2"
)

// NewLogger returns a new Logger
func NewLogger() *slog.Logger {
	return slog.New(slogzerolog.Option{Level: slog.LevelDebug}.NewZerologHandler())
}
