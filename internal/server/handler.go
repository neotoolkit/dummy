package server

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-dummy/dummy/internal/apischema"
	"github.com/go-dummy/dummy/internal/logger"
)

// Handler -.
type Handler struct {
	Path       string
	Method     string
	QueryParam url.Values
	Header     http.Header
	StatusCode int
	Response   interface{}
}

// Handlers -.
type Handlers struct {
	API    apischema.API
	Logger *logger.Logger
}

// NewHandlers returns a new instance of Handlers
func NewHandlers(api apischema.API, l *logger.Logger) Handlers {
	return Handlers{
		API:    api,
		Logger: l,
	}
}

// Get -.
func (h Handlers) Get(path, method string, queryParam url.Values, header http.Header, body io.ReadCloser) (apischema.Response, bool) {
	response, err := h.API.FindResponse(apischema.FindResponseParams{
		Path:   path,
		Method: method,
	})
	if err != nil {
		return apischema.Response{}, false
	}

	return response, true
}

// RemoveTrailingSlash returns path without trailing slash
func RemoveTrailingSlash(path string) string {
	if len(path) > 0 && path[len(path)-1] == '/' {
		return path[0 : len(path)-1]
	}

	return path
}

// RemoveFragment - clear url from reference in path
func RemoveFragment(path string) string {
	return RemoveTrailingSlash(strings.Split(path, "#")[0])
}
