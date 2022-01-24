package test_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lamoda/gonkey/runner"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/parse"
	"github.com/go-dummy/dummy/internal/server"
)

func TestDummy(t *testing.T) {
	testdata, err := ioutil.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}

	type test struct {
		SpecPath       string
		GonkeyTestsDir string
	}

	tests := make([]test, len(testdata))

	for i := 0; i < len(testdata); i++ {
		tests[i] = test{
			SpecPath:       "testdata/" + testdata[i].Name() + "/openapi.yml",
			GonkeyTestsDir: "testdata/" + testdata[i].Name() + "/cases",
		}
	}

	for _, tc := range tests {
		api, err := parse.Parse(tc.SpecPath)
		if err != nil {
			t.Fatal(err)
		}

		s := new(server.Server)
		conf := config.NewConfig()
		s.Config = conf.Server
		s.Handlers.API = api
		s.Logger = logger.NewLogger(conf.Logger.Level)

		mux := http.NewServeMux()
		mux.HandleFunc("/", s.Handler)
		newServer := httptest.NewServer(mux)

		runner.RunWithTesting(t, &runner.RunWithTestingParams{
			Server:   newServer,
			TestsDir: tc.GonkeyTestsDir,
		})
	}
}
