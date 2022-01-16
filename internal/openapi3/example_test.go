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

func TestExampleToResponse(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want interface{}
	}{
		{
			name: "",
			data: nil,
			want: nil,
		},
		{
			name: "",
			data: map[string]interface{}{},
			want: map[string]interface{}{},
		},
		{
			name: "",
			data: []interface{}{},
			want: []map[string]interface{}{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := openapi3.ExampleToResponse(tc.data)

			require.Equal(t, tc.want, got)
		})
	}
}
