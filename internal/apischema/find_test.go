package apischema_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/go-dummy/dummy/internal/apischema"
)

func FuzzPathByParamDetect(f *testing.F) {
	f.Add("", "", true)
	f.Add("/path", "/path", true)
	f.Add("/path/1", "/path/{1}", true)
	f.Add("/path/1/path/2", "/path/{1}/path/{2}", true)

	f.Fuzz(func(t *testing.T, path, param string, want bool) {
		got := apischema.PathByParamDetect(path, param)

		require.Equal(t, want, got)
	})
}
