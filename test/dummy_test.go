package test_test

import (
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
	api, err := parse.Parse("./testdata/openapi.yml")
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
		TestsDir: "./testdata/cases.yml",
	})
}
