package server_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/server"
)

func FuzzGetLastPathParam(f *testing.F) {
	f.Add("", "")
	f.Add("/path", "path")
	f.Add("/path/{path}", "{path}")

	f.Fuzz(func(t *testing.T, path, want string) {
		t.Parallel()

		got := server.GetLastPathSegment(path)

		if got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	})
}

func FuzzRemoveTrailingSlash(f *testing.F) {
	f.Add("", "")
	f.Add("/", "")
	f.Add("/path/", "/path")

	f.Fuzz(func(t *testing.T, path, want string) {
		t.Parallel()

		got := server.RemoveTrailingSlash(path)

		if got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	})
}

func FuzzIsLastPathSegmentParam(f *testing.F) {
	f.Add("", false)
	f.Add("/path", false)
	f.Add("/path/{path}", true)

	f.Fuzz(func(t *testing.T, path string, want bool) {
		t.Parallel()

		got := server.IsLastPathSegmentParam(path)

		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func FuzzParentPath(f *testing.F) {
	f.Add("", "")
	f.Add("/", "")
	f.Add("/path/", "/path")
	f.Add("/path/path", "/path")

	f.Fuzz(func(t *testing.T, path, want string) {
		t.Parallel()

		got := server.ParentPath(path)

		if got != want {
			t.Fatalf("got %s, want %s", got, want)
		}
	})
}

func FuzzPathByParamDetect(f *testing.F) {
	f.Add("", "", true)
	f.Add("/path", "/path", true)
	f.Add("/path/1", "/path/{1}", true)
	f.Add("/path/1/path/2", "/path/{1}/path/{2}", true)

	f.Fuzz(func(t *testing.T, path, param string, want bool) {
		t.Parallel()

		got := server.PathByParamDetect(path, param)

		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func FuzzRemoveFragment(f *testing.F) {
	f.Add("", "")
	f.Add("/", "")
	f.Add("/user", "/user")
	f.Add("/user#id", "/user")
	f.Add("/user#id,password", "/user")
	f.Add("/user/#id,password", "/user")
	f.Add("", "")
	f.Add("", "")

	f.Fuzz(func(t *testing.T, path, want string) {
		t.Parallel()

		got := server.RemoveFragment(path)

		if got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
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
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := server.RefSplit(tc.ref)

			require.Equal(t, tc.want, got)
		})
	}
}

func FuzzGetPathParamName(f *testing.F) {
	f.Add("", "")
	f.Add("{", "")
	f.Add("}", "")
	f.Add("{}", "")
	f.Add("some-string", "")
	f.Add("{some-string", "")
	f.Add("some-string}", "")
	f.Add("{some-string}", "some-string")

	f.Fuzz(func(t *testing.T, param, want string) {
		t.Parallel()

		got := server.GetPathParamName(param)

		if got != want {
			t.Fatalf(`got "%s", want "%s"`, got, want)
		}
	})
}
