package test_test

import (
	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDummy(t *testing.T) {
	testdata, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range testdata {
		path := "testdata/" + c.Name() + "/openapi3.yml"

		openapi, err := openapi3.Parse(path)
		if err != nil {
			t.Fatal(err)
		}

		s := new(server.Server)
		conf := config.NewConfig()
		s.Config = conf.Server
		s.OpenAPI = openapi
		s.Logger = logger.NewLogger()
		s.Handlers = make(map[string][]server.Handler)

		if err := s.SetHandlers(); err != nil {
			t.Fatal(err)
		}

		mux := http.NewServeMux()
		mux.HandleFunc("/", s.Handler)
		newServer := httptest.NewServer(mux)

		pact := dsl.Pact{
			Consumer:                 "consumer",
			Provider:                 "dummy",
			PactDir:                  "testdata/" + c.Name(),
			DisableToolValidityCheck: true,
		}

		if _, err := pact.VerifyProvider(t, types.VerifyRequest{
			ProviderBaseURL: newServer.URL,
			PactURLs:        []string{"testdata/" + c.Name() + "/pact.json"},
		}); err != nil {
			t.Fatal(err)
		}
	}
}
