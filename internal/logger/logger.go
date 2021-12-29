package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Logger is struct for Logger
type Logger struct {
	*zerolog.Logger
}

// NewLogger returns a new instance of Logger instance
func NewLogger(level string) *Logger {
	logger := zerolog.New(os.Stdout).Level(setLevel(level)).With().Timestamp().Logger()

	return &Logger{
		&logger,
	}
}

func setLevel(level string) zerolog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	default:
		return zerolog.InfoLevel
	}
}
