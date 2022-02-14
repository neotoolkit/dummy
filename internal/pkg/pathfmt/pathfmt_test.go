package pathfmt_test

import (
	"testing"

	"github.com/neotoolkit/dummy/internal/pkg/pathfmt"
	"github.com/stretchr/testify/require"
)

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
			got := pathfmt.RemoveTrailingSlash(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestRemoveFragment(t *testing.T) {
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
			path: "/#",
			want: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := pathfmt.RemoveFragment(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}
