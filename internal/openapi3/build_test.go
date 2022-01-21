package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestObjectExampleError(t *testing.T) {
	got := &openapi3.ObjectExampleError{
		Data: "",
	}

	require.Equal(t, got.Error(), "unpredicted type for example string")
}

func TestSchemaTypeError(t *testing.T) {
	got := &openapi3.SchemaTypeError{
		SchemaType: "",
	}

	require.Equal(t, got.Error(), "unknown type ")
}

func TestArrayExampleError(t *testing.T) {
	got := &openapi3.ArrayExampleError{
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
			got := openapi3.RemoveTrailingSlash(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}
