package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/api"
)

func TestObjectExampleError(t *testing.T) {
	got := &api.ObjectExampleError{
		Data: "",
	}

	require.Equal(t, got.Error(), "unpredicted type for example string")
}

func TestSchemaTypeError(t *testing.T) {
	got := &api.SchemaTypeError{
		SchemaType: "",
	}

	require.Equal(t, got.Error(), "unknown type ")
}

func TestArrayExampleError(t *testing.T) {
	got := &api.ArrayExampleError{
		Data: "",
	}

	require.Equal(t, got.Error(), "unpredicted type for example string")
}

func TestRemoveTrailingSlash(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "",
			path: "",
			want: "",
		},
		{
			name: "",
			path: "/",
			want: "",
		},
		{
			name: "",
			path: "path/",
			want: "path",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := api.RemoveTrailingSlash(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}
