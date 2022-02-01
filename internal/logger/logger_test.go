package logger_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neotoolkit/dummy/internal/logger"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name  string
		level string
	}{
		{
			name:  "",
			level: "",
		},
		{
			name:  "Logger with INFO level",
			level: "INFO",
		},
		{
			name:  "Logger with DEBUG level",
			level: "DEBUG",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := logger.NewLogger(tc.level)

			require.IsType(t, &logger.Logger{}, got)
		})
	}
}
