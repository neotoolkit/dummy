package server_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/server"
)

func TestGetLastPathParam(t *testing.T) {
	t.Parallel()

	type test struct {
		name string
		path string
		want string
	}

	tests := []test{
		{
			name: "",
			path: "",
			want: "",
		},
		{
			name: "",
			path: "/path",
			want: "path",
		},
		{
			name: "",
			path: "/path/{path}",
			want: "{path}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.GetLastPathSegment(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestRemoveTrailingSlash(t *testing.T) {
	t.Parallel()

	type test struct {
		name string
		path string
		want string
	}

	tests := []test{
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
			path: "/path/",
			want: "/path",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.RemoveTrailingSlash(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestIsLastPathSegmentParam(t *testing.T) {
	t.Parallel()

	type test struct {
		name string
		path string
		want bool
	}

	tests := []test{
		{
			name: "",
			path: "",
			want: false,
		},
		{
			name: "",
			path: "/path",
			want: false,
		},
		{
			name: "",
			path: "/path/{path}",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.IsLastPathSegmentParam(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestParentPath(t *testing.T) {
	t.Parallel()

	type test struct {
		name string
		path string
		want string
	}

	tests := []test{
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
			path: "/path/",
			want: "/path",
		},
		{
			name: "",
			path: "/path/path",
			want: "/path",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.ParentPath(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestPathByParamDetect(t *testing.T) {
	t.Parallel()

	type test struct {
		name  string
		path  string
		param string
		want  bool
	}

	tests := []test{
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
			path:  "/path/1/path/1",
			param: "/path/{1}/path",
			want:  false,
		},
		{
			name:  "",
			path:  "/path/1/path/1",
			param: "/path/{1}/path/{1}",
			want:  true,
		},
		{
			name:  "",
			path:  "/path/1/path/1",
			param: "/path/{1}/path/{1}",
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.PathByParamDetect(tc.path, tc.param)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestRemoveFragment(t *testing.T) {
	t.Parallel()

	type test struct {
		name string
		path string
		want string
	}

	tests := []test{
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
			path: "/user",
			want: "/user",
		},
		{
			name: "",
			path: "/user#id",
			want: "/user",
		},
		{
			name: "",
			path: "/user#id,password",
			want: "/user",
		},
		{
			name: "",
			path: "/user/#id,password",
			want: "/user",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.RemoveFragment(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestRefSplit(t *testing.T) {
	t.Parallel()

	type test struct {
		name string
		ref  string
		want []string
	}

	tests := []test{
		{
			name: "",
			ref:  "",
			want: nil,
		},
		{
			name: "",
			ref:  "/",
			want: nil,
		},
		{
			name: "",
			ref:  "#/",
			want: nil,
		},
		{
			name: "",
			ref:  "#/components",
			want: []string{"components"},
		},
		{
			name: "",
			ref:  "#/components/schemas/User",
			want: []string{"components", "schemas", "User"},
		},
		{
			name: "",
			ref:  "#/components/schemas/Users",
			want: []string{"components", "schemas", "Users"},
		},
		{
			name: "",
			ref:  "#/components/schemas/Users/",
			want: []string{"components", "schemas", "Users"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.RefSplit(tc.ref)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestEqualHeadersByValues(t *testing.T) {
	t.Parallel()

	type test struct {
		name string
		h1   []string
		h2   []string
		want bool
	}

	tests := []test{
		{
			name: "",
			h1:   nil,
			h2:   nil,
			want: true,
		},
		{
			name: "",
			h1:   []string{"1"},
			h2:   nil,
			want: false,
		},
		{
			name: "",
			h1:   nil,
			h2:   []string{"1"},
			want: false,
		},
		{
			name: "",
			h1:   []string{"1"},
			h2:   []string{"1"},
			want: true,
		},
		{
			name: "",
			h1:   []string{"1"},
			h2:   []string{"1", "2"},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.EqualHeadersByValues(tc.h1, tc.h2)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestGetPathParamName(t *testing.T) {
	type test struct {
		input, result string
	}

	tests := []test{
		{"{some-string}", "some-string"},
		{"{some-string", ""},
		{"some-string}", ""},
		{"some-string", ""},
		{"", ""},
		{"{", ""},
		{"}", ""},
	}

	for _, tc := range tests {
		p := server.GetPathParamName(tc.input)

		assert.Equal(t, tc.result, p)
	}
}
