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

func TestLastPathSegmentIsParam(t *testing.T) {
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
		got := server.LastPathSegmentIsParam(tests[i].path)
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
		got := server.PathByParamDetect(tests[i].path, tests[i].mask)
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
		{"/", ""},
		{"/user", "/user"},
		{"/user#id", "/user"},
		{"/user#id,password", "/user"},
		{"/user/#id,password", "/user"},
	}

	for _, tc := range tests {
		p := server.RemoveFragment(tc.request)

		assert.Equal(t, tc.response, p)
	}
}

func TestRefSplit(t *testing.T) {
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

	for _, tc := range tests {
		p := server.RefSplit(tc.request)

		assert.Equal(t, tc.response, p)
	}
}

func TestEqualHeadersByValues(t *testing.T) {
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

	for _, tc := range tests {
		p := server.EqualHeadersByValues(tc.h1, tc.h2)

		assert.Equal(t, tc.result, p)
	}
}

func TestGetPathParamName(t *testing.T) {
	type test struct {
		input, result string
	}

	tests := []test{
		{"{some string}", "some string"},
		{"{some string", ""},
		{"some string}", ""},
		{"some string", ""},
	}

	for _, tc := range tests {
		p := server.GetPathParamName(tc.input)
		assert.Equal(t, tc.result, p)
	}
}
