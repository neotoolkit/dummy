package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestExamplesGetKeys(t *testing.T) {
	e := openapi3.Examples{
		"first_example":  {},
		"second_example": {},
	}

	res := e.GetKeys()

	require.Equal(t, len(e), len(res))
}
