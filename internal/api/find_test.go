package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/api"
)

func TestPathByParamDetect(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		param string
		want  bool
	}{
		{
			name:  "",
			path:  "",
			param: "",
			want:  true,
		},
		{
			name:  "",
			path:  "/path",
			param: "/path",
			want:  true,
		},
		{
			name:  "",
			path:  "/path/1",
			param: "/path/{1}",
			want:  true,
		},
		{
			name:  "",
			path:  "/path/1/path/2",
			param: "/path/{1}/path/{2}",
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := api.PathByParamDetect(tc.path, tc.param)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestFindResponseError(t *testing.T) {
	got := &api.FindResponseError{
		Method: "test method",
		Path:   "test path",
	}

	require.Equal(t, got.Error(), "not specified operation: test method test path")
}
