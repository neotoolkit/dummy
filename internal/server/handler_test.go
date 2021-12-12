package server_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/server"
)

func TestGetLastPathParam(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path string
		want string
	}{
		{
			path: "",
			want: "",
		},
		{
			path: "/path",
			want: "path",
		},
		{
			path: "/path/{path}",
			want: "{path}",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := server.GetLastPathSegment(tt.path)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestRemoveTrailingSlash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path string
		want string
	}{
		{
			path: "",
			want: "",
		},
		{
			path: "/",
			want: "",
		},
		{
			path: "/path/",
			want: "/path",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := server.RemoveTrailingSlash(tt.path)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestLastPathSegmentIsParam(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path string
		want bool
	}{
		{
			path: "",
			want: false,
		},
		{
			path: "/path",
			want: false,
		},
		{
			path: "/path/{path}",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := server.LastPathSegmentIsParam(tt.path)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestParentPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path string
		want string
	}{
		{
			path: "",
			want: "",
		},
		{
			path: "/",
			want: "",
		},
		{
			path: "/path/",
			want: "/path",
		},
		{
			path: "/path/path",
			want: "/path",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := server.ParentPath(tt.path)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestPathByMaskDetect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path string
		mask string
		want bool
	}{
		{
			path: "",
			mask: "",
			want: true,
		},
		{
			path: "/path",
			mask: "/path",
			want: true,
		},
		{
			path: "/path/1",
			mask: "/path/{1}",
			want: true,
		},
		{
			path: "/path/1/path/1",
			mask: "/path/{1}/path",
			want: false,
		},
		{
			path: "/path/1/path/1",
			mask: "/path/{1}/path/{1}",
			want: true,
		},
		{
			path: "/path/1/path/1",
			mask: "/path/{1}/path/{1}",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := server.PathByParamDetect(tt.path, tt.mask)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestRemoveFragment(t *testing.T) {
	t.Parallel()

	type test struct {
		request  string
		response string
	}

	tests := []test{
		{"", ""},
		{"/", ""},
		{"/user", "/user"},
		{"/user#id", "/user"},
		{"/user#id,password", "/user"},
		{"/user/#id,password", "/user"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			p := server.RemoveFragment(tt.request)
			require.Equal(t, tt.response, p)
		})
	}
}

func TestRefSplit(t *testing.T) {
	t.Parallel()

	type test struct {
		request  string
		response []string
	}

	tests := []test{
		{"", nil},
		{"/", nil},
		{"#/", nil},
		{"#/components", []string{"components"}},
		{"#/components/schemas/User", []string{"components", "schemas", "User"}},
		{"#/components/schemas/Users", []string{"components", "schemas", "Users"}},
		{"#/components/schemas/Users/", []string{"components", "schemas", "Users"}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			p := server.RefSplit(tt.request)
			require.Equal(t, tt.response, p)
		})
	}
}

func TestEqualHeadersByValues(t *testing.T) {
	t.Parallel()

	type test struct {
		h1     []string
		h2     []string
		result bool
	}

	tests := []test{
		{nil, nil, true},
		{[]string{"1"}, nil, false},
		{nil, []string{"1"}, false},
		{[]string{"1"}, []string{"1"}, true},
		{[]string{"1"}, []string{"1", "2"}, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			p := server.EqualHeadersByValues(tt.h1, tt.h2)
			require.Equal(t, tt.result, p)
		})
	}
}
