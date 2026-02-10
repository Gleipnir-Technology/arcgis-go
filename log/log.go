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

func Debug() *zerolog.Event {
	return Logger.Debug()
}
func Error() *zerolog.Event {
	return Logger.Error()
}
// LoggerFromContext gets the logger from context or falls back to the default
func FromContext(ctx context.Context) zerolog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(zerolog.Logger); ok {
		return logger
	}
	return Logger // Fall back to package-level logger
}

func Info() *zerolog.Event {
	return Logger.Info()
}
// SetLogLevel allows clients to control this library's log verbosity
func SetLogLevel(level zerolog.Level) {
	Logger = Logger.Level(level)
}

// SetLogOutput allows clients to redirect logs if needed
func SetLogOutput(w io.Writer) {
	Logger = zerolog.New(w).With().Timestamp().Logger().Level(Logger.GetLevel())
}

func Warn() *zerolog.Event {
	return Logger.Warn()
}
// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}
