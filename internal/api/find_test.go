//go:build go1.18
// +build go1.18

package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/api"
)

func FuzzPathByParamDetect(f *testing.F) {
	tests := []struct {
		path  string
		param string
		want  bool
	}{
		{
			path:  "",
			param: "",
			want:  true,
		},
		{
			path:  "/path",
			param: "/path",
			want:  true,
		},
		{
			path:  "/path/1",
			param: "/path/{1}",
			want:  true,
		},
		{
			path:  "/path/1/path/2",
			param: "/path/{1}/path/{2}",
			want:  true,
		},
	}

	for _, seed := range tests {
		f.Add(seed.path, seed.param, seed.want)
	}

	f.Fuzz(func(t *testing.T, path, param string, want bool) {
		got := api.PathByParamDetect(path, param)

		require.Equal(t, want, got)
	})
}
