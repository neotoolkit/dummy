package logger

import (
	"github.com/rs/zerolog"
	"os"
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
