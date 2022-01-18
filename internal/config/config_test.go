package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/config"
)

func TestNewConfig(t *testing.T) {
	conf := config.NewConfig()

	require.IsType(t, &config.Config{}, conf)
}
