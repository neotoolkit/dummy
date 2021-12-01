package server_test

import (
	"testing"

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
