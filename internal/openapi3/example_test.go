package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestExamples_GetKeys(t *testing.T) {
	e := openapi3.Examples{
		"first_example":  openapi3.Example{},
		"second_example": openapi3.Example{},
	}

	res := e.GetKeys()

	require.Equal(t, len(e), len(res))
}
