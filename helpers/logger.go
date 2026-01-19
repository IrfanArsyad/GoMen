package helpers

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger zerolog.Logger

// InitLogger initializes the global logger with appropriate settings
func InitLogger(debug bool, appEnv string) {
	// Set global log level
	zerolog.TimeFieldFormat = time.RFC3339

	var output io.Writer = os.Stdout

	// Pretty logging for development
	if debug || appEnv == "development" {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
		}
	}

	Logger = zerolog.New(output).
		With().
		Timestamp().
		Caller().
		Logger()

	// Set as global logger
	log.Logger = Logger

	// Set log level
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// GetLogger returns the global logger instance
func GetLogger() *zerolog.Logger {
	return &Logger
}

// Info logs an info level message
func Info(msg string) *zerolog.Event {
	return Logger.Info().Str("level", "info")
}

// Error logs an error level message
func Error(err error, msg string) *zerolog.Event {
	return Logger.Error().Err(err).Str("level", "error")
}

// Debug logs a debug level message
func Debug(msg string) *zerolog.Event {
	return Logger.Debug().Str("level", "debug")
}

// Warn logs a warning level message
func Warn(msg string) *zerolog.Event {
	return Logger.Warn().Str("level", "warn")
}

// Fatal logs a fatal level message and exits
func Fatal(err error, msg string) *zerolog.Event {
	return Logger.Fatal().Err(err).Str("level", "fatal")
}
