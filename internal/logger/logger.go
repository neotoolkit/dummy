package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func NewLogger() *Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{
		&logger,
	}
}
