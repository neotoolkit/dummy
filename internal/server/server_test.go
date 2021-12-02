package server_test

import (
	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewServerRunError(t *testing.T) {
	s := server.NewServer(config.Server{Port: "test"}, openapi3.OpenAPI{})

	require.EqualError(t, s.Run(), "listen tcp: lookup tcp/test: nodename nor servname provided, or not known")
}
