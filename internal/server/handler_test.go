package server_test

import (
	"github.com/neotoolkit/dummy/internal/api"
	"github.com/neotoolkit/dummy/internal/logger"
	"github.com/neotoolkit/dummy/internal/server"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHandlers(t *testing.T) {
	h := server.NewHandlers(api.API{}, &logger.Logger{})

	require.IsType(t, server.Handlers{}, h)
}
