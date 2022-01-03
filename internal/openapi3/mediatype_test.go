package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestMediaType_ResponseByExample(t *testing.T) {
	m := openapi3.MediaType{
		Example: []interface{}{},
	}

	require.IsType(t, []map[string]interface{}{}, m.ResponseByExample())
}

func TestMediaType_ResponseByExamplesKey(t *testing.T) {
	const key = "key"

	m := openapi3.MediaType{
		Examples: openapi3.Examples{
			key: openapi3.Example{
				Value: map[string]interface{}{
					"key": "value",
				},
			},
		},
	}

	require.IsType(t, map[string]interface{}{"key": "value"}, m.ResponseByExamplesKey(key))
}
