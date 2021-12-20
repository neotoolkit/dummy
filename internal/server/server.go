package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-dummy/dummy/internal/config"
	"github.com/go-dummy/dummy/internal/logger"
)

// Server is struct for Server
type Server struct {
	Config   config.Server
	Server   *http.Server
	Logger   *logger.Logger
	Handlers Handlers
}

// NewServer returns a new instance of Server instance
func NewServer(config config.Server, l *logger.Logger, h Handlers) *Server {
	return &Server{
		Config:   config,
		Logger:   l,
		Handlers: h,
	}
}

// Run -.
func (s *Server) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.Handler)

	s.Server = &http.Server{
		Addr:    ":" + s.Config.Port,
		Handler: mux,
	}

	return s.Server.ListenAndServe()
}

// Handler -.
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Set-Status-Code") == "500" {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if response, ok := s.Handlers.Get(RemoveFragment(r.URL.Path), r.Method, r.URL.Query(), r.Header, r.Body); ok {
		w.WriteHeader(response.StatusCode)
		bytes, err := json.Marshal(response.ExampleValue(r.Header.Get("X-Example")))
		if err != nil {
			s.Logger.Error().Err(err).Msg("serialize response")
		}
		_, err = w.Write(bytes)
		if err != nil {
			s.Logger.Error().Err(err).Msg("write response")
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)
}
