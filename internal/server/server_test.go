package server_test

import (
	"context"
	"errors"
	"github.com/neotoolkit/dummy/internal/config"
	"github.com/neotoolkit/dummy/internal/logger"
	"github.com/neotoolkit/dummy/internal/server"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s := server.NewServer(config.Server{}, &logger.Logger{}, server.Handlers{})

	require.IsType(t, &server.Server{}, s)
}

func TestNewServer_Run(t *testing.T) {
	l := logger.NewLogger("DEBUG")

	s := server.NewServer(config.Server{}, l, server.Handlers{})

	go func() {
		time.Sleep(1 * time.Second)
		s.Stop(context.Background())
	}()

	err := s.Run()

	require.True(t, errors.Is(err, http.ErrServerClosed))
}
