package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
)

func TestCheck(t *testing.T) {
	cases, err := ioutil.ReadDir("./cases")
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		path := "cases/" + c.Name() + "/openapi3.yaml"

		openapi, err := openapi3.Parse(path)
		if err != nil {
			t.Fatal(err)
		}

		s := new(server.Server)
		s.Config = config.Server{
			Port: "4000",
			Path: path,
		}
		s.OpenAPI = openapi
		s.Logger = logger.NewLogger()
		s.Handlers = make(map[string]server.Handler)

		if err := s.SetHandlers(); err != nil {
			t.Fatal(err)
		}

		mux := http.NewServeMux()
		mux.HandleFunc("/", s.Handler)
		newServer := httptest.NewServer(mux)

		for k, v := range s.OpenAPI.Paths {
			t.Run(c.Name(), func(t *testing.T) {
				if v.Post != nil {
					makeTestReq(t, http.MethodPost, newServer.URL+k, c.Name())
				}
				if v.Get != nil {
					makeTestReq(t, http.MethodGet, newServer.URL+k, c.Name())
				}
			})
		}
	}
}

func makeTestReq(t *testing.T, method, url, testCase string) {
	t.Helper()

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	h, err := ioutil.ReadFile("cases/" + testCase + "/header.txt")
	if err != nil {
		t.Fatal(err)
	}

	if len(h) > 0 {
		headers := strings.Split(string(h), `\n`)

		for i := 0; i < len(headers); i++ {
			header := strings.Split(headers[i], ":")
			key := header[0]
			value := strings.TrimSpace(header[1])

			req.Header.Set(key, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	r, err := ioutil.ReadFile("cases/" + testCase + "/" + method + ".response.json")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNotFound {
		require.JSONEq(t, string(out), string(r))
	}
}
