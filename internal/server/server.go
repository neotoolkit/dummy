package server

import (
	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/openapi3"
	"net/http"
)

type Server struct {
	cfg     config.Server
	openapi openapi3.OpenAPI
}

func NewServer(cfg config.Server, openapi openapi3.OpenAPI) *Server {
	return &Server{
		cfg:     cfg,
		openapi: openapi,
	}
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.Handler)

	return http.ListenAndServe(":"+s.cfg.Port, nil)
}
