package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

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
		path := "cases/" + c.Name() + "/openapi3.yml"

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
		s.Handlers = make(map[string][]server.Handler)

		if err := s.SetHandlers(); err != nil {
			t.Fatal(err)
		}

		mux := http.NewServeMux()
		mux.HandleFunc("/", s.Handler)
		newServer := httptest.NewServer(mux)

		for key, data := range s.Handlers {
			for i := 0; i < len(data); i++ {
				t.Run(c.Name(), func(t *testing.T) {
					makeTestReq(t, data[i].Method, key, newServer.URL+data[i].Path, c.Name())
				})
			}
		}
	}
}

func makeTestReq(t *testing.T, method, path, url, testCase string) {
	t.Helper()

	req, err := http.NewRequest(method, changePathMask(url), nil)
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
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	p := strings.Split(path, "/")
	newPath := strings.Join(p[1:], "-")

	responses, err := ioutil.ReadDir("cases/" + testCase + "/" + newPath + "/" + method)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(responses); i++ {
		r, err := ioutil.ReadFile("cases/" + testCase + "/" + newPath + "/" + method + "/" + responses[i].Name())
		if err != nil {
			t.Fatal(err)
		}

		equal, err := jsonBytesEqual(out, r)
		if err != nil {
			t.Fatal(err)
		}

		if equal {
			return
		}
	}

	t.Fatal()
}

func jsonBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}

	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}

	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(j2, j), nil
}

func changePathMask(path string) string {
	p := strings.Split(path, "/")

	for i := 0; i < len(p); i++ {
		if strings.HasPrefix(p[i], "{") && strings.HasSuffix(p[i], "}") {
			p[i] = "1"
		}
	}

	return strings.Join(p, "/")
}
