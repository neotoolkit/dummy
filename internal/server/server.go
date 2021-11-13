package server

import (
	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/openapi3"
	"net/http"
)

type Server struct {
	cfg      config.Server
	openapi  openapi3.OpenAPI
	handlers map[string]Handler
}

func NewServer(cfg config.Server, openapi openapi3.OpenAPI) *Server {
	return &Server{
		cfg:      cfg,
		openapi:  openapi,
		handlers: make(map[string]Handler, len(openapi.Paths)),
	}
}

func (s *Server) Run() error {
	if err := s.Handlers(); err != nil {
		return err
	}

	http.HandleFunc("/", s.Handler)

	return http.ListenAndServe(":"+s.cfg.Port, nil)
}
