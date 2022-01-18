package openapi3_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestSchemaError(t *testing.T) {
	got := &openapi3.SchemaError{
		Ref: "test",
	}

	require.Equal(t, got.Error(), "unknown schema test")
}

func TestLookupByReference(t *testing.T) {
	api := openapi3.OpenAPI{}

	schema, err := api.LookupByReference("")

	var schemaErr *openapi3.SchemaError

	require.Equal(t, openapi3.Schema{}, schema)
	require.True(t, errors.As(err, &schemaErr))
}
