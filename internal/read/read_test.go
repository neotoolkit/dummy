package read_test

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/read"
)

func TestRead(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{
   "openapi": "3.0.3",
   "info": {
      "title": "Test dummy API",
      "version": "0.1.0"
   },
   "paths": {
      "/healthz": {
         "get": {
            "responses": {
               "201": {
                  "description": ""
               }
            }
         }
      }
   }
}`)
	})
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Length", "50")
		// return less bytes, which will result in an "unexpected EOF" from ioutil.ReadAll()
		fmt.Fprintln(w, []byte("a"))
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	tests := []struct {
		name string
		path string
		err  error
	}{
		{
			name: "empty URL",
			path: "https://",
			err: &url.Error{
				Op:  "Get",
				URL: "https:",
				Err: errors.New("http: no Host in request URL"),
			},
		},
		{
			name: "read from URL error",
			path: ts.URL + "/error",
			err:  errors.New("unexpected EOF"),
		},
		{
			name: "read from URL",
			path: ts.URL,
			err:  nil,
		},
		{
			name: "read from file",
			path: "./testdata/openapi3.yml",
			err:  nil,
		},
		{
			name: "empty file",
			path: "",
			err: &fs.PathError{
				Op:   "open",
				Path: "",
				Err:  errors.New("no such file or directory"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := read.Read(tc.path)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}

			require.IsType(t, []byte{}, got)
		})
	}
}
