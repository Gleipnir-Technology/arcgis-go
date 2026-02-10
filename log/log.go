package log

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

type loggerKey struct{}

// SetLogLevel allows clients to control this library's log verbosity
func SetLogLevel(level zerolog.Level) {
	Logger = Logger.Level(level)
}

// SetLogOutput allows clients to redirect logs if needed
func SetLogOutput(w io.Writer) {
	Logger = zerolog.New(w).With().Timestamp().Logger().Level(Logger.GetLevel())
}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// LoggerFromContext gets the logger from context or falls back to the default
func LoggerFromContext(ctx context.Context) zerolog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(zerolog.Logger); ok {
		return logger
	}
	return Logger // Fall back to package-level logger
}

func DoSomethingWithContext(ctx context.Context) {
	logger := LoggerFromContext(ctx)
	logger.Debug().Msg("Debug from context-aware function")
	// ...
}

