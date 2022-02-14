package server_test

import (
	"github.com/neotoolkit/dummy/internal/config"
	"github.com/neotoolkit/dummy/internal/logger"
	"github.com/neotoolkit/dummy/internal/server"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := server.NewServer(config.Server{}, &logger.Logger{}, server.Handlers{})

	require.IsType(t, &server.Server{}, s)
}
