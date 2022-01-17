package openapi3_test

import (
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchemaError(t *testing.T) {
	got := &openapi3.SchemaError{
		Ref: "test",
	}

	require.Equal(t, got.Error(), "unknown schema test")
}
