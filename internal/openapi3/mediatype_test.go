package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestMediaTypeResponseByExample(t *testing.T) {
	m := openapi3.MediaType{
		Example: []interface{}{},
	}

	require.IsType(t, []map[string]interface{}{}, m.ResponseByExample())
}

func TestMediaTypeResponseByExamplesKey(t *testing.T) {
	const key = "key"

	m := openapi3.MediaType{
		Examples: openapi3.Examples{
			key: openapi3.Example{
				Value: map[interface{}]interface{}{
					"key": "value",
				},
			},
		},
	}

	require.IsType(t, map[string]interface{}{"key": "value"}, m.ResponseByExamplesKey(key))
}
