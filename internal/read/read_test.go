package read_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/read"
)

func TestRead(t *testing.T) {
	type test struct {
		name string
		path string
		err  error
	}

	tests := []test{
		{
			name: "read file",
			path: "./testdata/openapi3.yml",
			err:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := read.Read(tc.path)

			require.IsType(t, []byte{}, got)
			require.True(t, errors.Is(err, tc.err))
		})
	}
}
