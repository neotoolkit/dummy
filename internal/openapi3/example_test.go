package openapi3_test

import (
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExamplesGetExamplesKeys(t *testing.T) {
	e := openapi3.Examples{
		"first_example":  {},
		"second_example": {},
	}

	keys := []string{"first_example", "second_example"}
	res := e.GetExamplesKeys()

	require.Equal(t, len(e), len(res))
	require.Equal(t, keys, res)
}
