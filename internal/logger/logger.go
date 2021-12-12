package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger is struct for Logger
type Logger struct {
	*zerolog.Logger
}

// NewLogger returns a new instance of Logger instance
func NewLogger() *Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{
		&logger,
	}
}
