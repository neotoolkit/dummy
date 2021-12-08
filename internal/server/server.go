package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
)

type Server struct {
	Config   config.Server
	Server   *http.Server
	Logger   *logger.Logger
	Handlers Handlers
}

func NewServer(config config.Server, openapi openapi3.OpenAPI) *Server {
	return &Server{
		Config: config,
		Logger: logger.NewLogger(),
		Handlers: Handlers{
			OpenAPI:  openapi,
			Handlers: make(map[string][]Handler),
		},
	}
}

func (s *Server) Run() error {
	if err := s.Handlers.Init(); err != nil {
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

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Set-Status-Code") == "500" {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if h, ok := s.Handlers.Get(RemoveFragment(r.URL.Path), r.Method, r.URL.Query(), r.Header, r.Body); ok {
		w.WriteHeader(h.StatusCode)
		bytes, _ := json.Marshal(h.Response)
		_, _ = w.Write(bytes)

		return
	}

	w.WriteHeader(http.StatusNotFound)
}
