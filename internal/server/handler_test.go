package server_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-dummy/dummy/internal/server"
)

func TestGetLastPathParam(t *testing.T) {
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

	for i := 0; i < len(tests); i++ {
		got := server.GetLastPathSegment(tests[i].path)
		if tests[i].want != got {
			t.Fatalf(`expected: "%s", got: "%s"`, tests[i].want, got)
		}
	}
}

func TestRemoveTrailingSlash(t *testing.T) {
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

	for i := 0; i < len(tests); i++ {
		got := server.RemoveTrailingSlash(tests[i].path)
		if tests[i].want != got {
			t.Fatalf(`expected: "%s", got: "%s"`, tests[i].want, got)
		}
	}
}

func TestLastParamIsMask(t *testing.T) {
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

	for i := 0; i < len(tests); i++ {
		got := server.LastParamIsMask(tests[i].path)
		if tests[i].want != got {
			t.Fatalf(`expected: "%v", got: "%v"`, tests[i].want, got)
		}
	}
}

func TestParentPath(t *testing.T) {
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

	for i := 0; i < len(tests); i++ {
		got := server.ParentPath(tests[i].path)
		if tests[i].want != got {
			t.Fatalf(`expected: "%s", got: "%s"`, tests[i].want, got)
		}
	}
}

func TestPathByMaskDetect(t *testing.T) {
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

	for i := 0; i < len(tests); i++ {
		got := server.PathByMaskDetect(tests[i].path, tests[i].mask)
		if tests[i].want != got {
			t.Fatalf(`expected: "%v", got: "%v"`, tests[i].want, got)
		}
	}
}

func TestRemoveFragment(t *testing.T) {
	type test struct {
		request  string
		response string
	}

	tests := []test{
		{"", ""},
		{"/", "/"},
		{"/user", "/user"},
		{"/user#id", "/user"},
		{"/user#id,password", "/user"},
	}

	for _, tc := range tests {
		p := server.RemoveFragment(tc.request)

		assert.Equal(t, tc.response, p)
	}
}
