package logger_test

import (
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewLogger(t *testing.T) {
	l := logger.NewLogger()

	require.IsType(t, &logger.Logger{}, l)
}
