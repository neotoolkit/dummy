package logger_test

import (
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLogger(t *testing.T) {
	type test struct {
		name  string
		level string
	}

	tests := []test{
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