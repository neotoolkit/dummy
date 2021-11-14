package server

import (
	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/openapi3"
	"net/http"
)

type Server struct {
	Cfg      config.Server
	OpenAPI  openapi3.OpenAPI
	Handlers map[string]Handler
	Server   *http.Server
}

func NewServer(cfg config.Server, openapi openapi3.OpenAPI) *Server {
	return &Server{
		Cfg:      cfg,
		OpenAPI:  openapi,
		Handlers: make(map[string]Handler, len(openapi.Paths)),
	}
}

func (s *Server) Run() error {
	if err := s.SetHandlers(); err != nil {
		return err
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", s.Handler)

	s.Server = &http.Server{
		Addr:    ":" + s.Cfg.Port,
		Handler: mux,
	}

	return s.Server.ListenAndServe()
}
