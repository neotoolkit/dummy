package server

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/neotoolkit/dummy/internal/api"
	"github.com/neotoolkit/dummy/internal/logger"
	"github.com/neotoolkit/dummy/internal/pkg/pathfmt"
)

// Handlers -.
type Handlers struct {
	API    api.API
	Logger *logger.Logger
}

// NewHandlers returns a new instance of Handlers
func NewHandlers(api api.API, l *logger.Logger) Handlers {
	return Handlers{
		API:    api,
		Logger: l,
	}
}

// Handler -.
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	if setStatusCode(w, r.Header.Get("X-Set-Status-Code")) {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	p := pathfmt.RemoveFragment(r.URL.Path)

	response, ok, err := s.Handlers.Get(p, r.Method, r.Body)
	if ok {
		if _, ok := err.(*json.SyntaxError); ok || errors.Is(err, api.ErrEmptyRequireField) || err == io.EOF {
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		w.WriteHeader(response.StatusCode)
		resp := response.ExampleValue(r.Header.Get("X-Example"))

		if nil == resp {
			return
		}

		bytes, err := json.Marshal(resp)
		if err != nil {
			s.Logger.Error().Err(err).Msg("serialize response")
		}

		if _, err := w.Write(bytes); err != nil {
			s.Logger.Error().Err(err).Msg("write response")
		}

		return
	}

	w.WriteHeader(http.StatusNotFound)
}

// Get -.
func (h Handlers) Get(path, method string, body io.ReadCloser) (api.Response, bool, error) {
	response, err := h.API.FindResponse(api.FindResponseParams{
		Path:   path,
		Method: method,
		Body:   body,
	})
	if err != nil {
		if errors.Is(err, api.ErrEmptyRequireField) {
			return api.Response{}, true, err
		}

		if err == io.EOF {
			return api.Response{}, true, err
		}

		if _, ok := err.(*json.SyntaxError); ok {
			return api.Response{}, true, err
		}

		return api.Response{}, false, err
	}

	return response, true, nil
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
