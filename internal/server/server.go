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
	server   *http.Server
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

	mux := http.NewServeMux()

	mux.HandleFunc("/", s.Handler)

	s.server = &http.Server{
		Addr:    ":" + s.cfg.Port,
		Handler: mux,
	}

	return s.server.ListenAndServe()
}
