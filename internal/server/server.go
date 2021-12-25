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

	s.Logger.Info().Msgf("Running mock server on %s port", s.Config.Port)

	err := s.Server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// Handler -.
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	if setStatusCode(w, r.Header.Get("X-Set-Status-Code")) {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	path := RemoveFragment(r.URL.Path)

	if response, ok := s.Handlers.Get(path, r.Method, r.URL.Query(), r.Header, r.Body); ok {
		w.WriteHeader(response.StatusCode)
		resp := response.ExampleValue(r.Header.Get("X-Example"))
		if resp == nil {
			return
		}
		bytes, err := json.Marshal(resp)
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

func setStatusCode(w http.ResponseWriter, statusCode string) bool {
	switch statusCode {
	case "500":
		w.WriteHeader(http.StatusInternalServerError)

		return true
	default:
		return false
	}
}
