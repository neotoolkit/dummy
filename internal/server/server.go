package server

import (
	"net/http"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
)

type Server struct {
	Config   config.Server
	OpenAPI  openapi3.OpenAPI
	Server   *http.Server
	Logger   *logger.Logger
	Handlers map[string][]Handler
}

func NewServer(config config.Server, openapi openapi3.OpenAPI) *Server {
	return &Server{
		Config:   config,
		OpenAPI:  openapi,
		Logger:   logger.NewLogger(),
		Handlers: make(map[string][]Handler),
	}
}

func (s *Server) Run() error {
	if err := s.SetHandlers(); err != nil {
		return err
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", s.Handler)

	s.Server = &http.Server{
		Addr:    ":" + s.Config.Port,
		Handler: mux,
	}

	return s.Server.ListenAndServe()
}
