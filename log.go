package arcgis

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// No default logger declaration - use the global one instead

// For context-based logging
type loggerKey struct{}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// LoggerFromContext gets the logger from context or falls back to the global logger
func LoggerFromContext(ctx context.Context) zerolog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(zerolog.Logger); ok {
		return logger
	}
	return log.Logger // Fall back to global logger (uses client's setup)
}

// Helper to create a library-specific logger with consistent context
func libraryLogger(ctx context.Context) zerolog.Logger {
	return LoggerFromContext(ctx).With().Str("component", "mylibrary").Logger()
}
