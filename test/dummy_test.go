package test_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
	"github.com/go-dummy/dummy/internal/server"
)

func FuzzDummy(f *testing.F) {
	testdata, err := ioutil.ReadDir("testdata")
	if err != nil {
		f.Fatal(err)
	}

	for i := 0; i < len(testdata); i++ {
		f.Add(
			"testdata/"+testdata[i].Name()+"/openapi3.yml",
			"testdata/"+testdata[i].Name(),
			"testdata/"+testdata[i].Name()+"/pact.json",
		)
	}

	f.Fuzz(func(t *testing.T, path, pactDir, pactURL string) {
		api, err := openapi3.Parse(path)
		if err != nil {
			t.Fatal(err)
		}

		s := new(server.Server)
		conf := config.NewConfig()
		s.Config = conf.Server
		s.Handlers.API = api
		s.Logger = logger.NewLogger()
		s.Handlers.Handlers = make(map[string][]server.Handler)

		mux := http.NewServeMux()
		mux.HandleFunc("/", s.Handler)
		newServer := httptest.NewServer(mux)

		pact := dsl.Pact{
			Consumer:                 "consumer",
			Provider:                 "dummy",
			PactDir:                  pactDir,
			DisableToolValidityCheck: true,
		}

		if _, err := pact.VerifyProvider(t, types.VerifyRequest{
			ProviderBaseURL: newServer.URL,
			PactURLs:        []string{pactURL},
		}); err != nil {
			t.Fatal(err)
		}
	})
}
