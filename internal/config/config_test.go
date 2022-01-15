package config_test

import (
	"github.com/go-dummy/dummy/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConfig(t *testing.T) {
	conf := config.NewConfig()

	require.IsType(t, &config.Config{}, conf)
}
