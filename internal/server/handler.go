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
	Response   any
}

// Handlers -.
type Handlers struct {
	API      apischema.API
	Handlers map[string][]Handler
	Logger   *logger.Logger
}

// NewHandlers returns a new instance of Handlers
func NewHandlers(api apischema.API, l *logger.Logger) Handlers {
	return Handlers{
		API:      api,
		Handlers: make(map[string][]Handler),
		Logger:   l,
	}
}

// Get -.
func (h Handlers) Get(path, method string, queryParam url.Values, header http.Header, body io.ReadCloser) (apischema.Response, bool) {
	response, err := h.API.FindResponse(apischema.FindResponseParams{
		Path:   path,
		Method: method,
		//MediaType: header.Get("Accept"),
	})
	if err != nil {
		return apischema.Response{}, false
	}

	return response, true
}

// PathByParamDetect returns result of
func PathByParamDetect(path, param string) bool {
	splitPath := strings.Split(path, "/")
	splitParam := strings.Split(param, "/")

	if len(splitPath) != len(splitParam) {
		return false
	}

	for i := 0; i < len(splitPath); i++ {
		if strings.HasPrefix(splitParam[i], "{") && strings.HasSuffix(splitParam[i], "}") {
			continue
		}

		if splitPath[i] != splitParam[i] {
			return false
		}
	}

	return true
}

// ParentPath returns parent path
func ParentPath(path string) string {
	p := strings.Split(path, "/")

	return strings.Join(p[0:len(p)-1], "/")
}

// IsLastPathSegmentParam returns the result of checking whether last path segment is a param
func IsLastPathSegmentParam(path string) bool {
	p := strings.Split(path, "/")

	return strings.HasPrefix(p[len(p)-1], "{") && strings.HasSuffix(p[len(p)-1], "}")
}

// GetLastPathSegment returns last path segment
func GetLastPathSegment(path string) string {
	p := strings.Split(path, "/")

	return p[len(p)-1]
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

// RefSplit - returns a list of references
func RefSplit(ref string) []string {
	if len(ref) > 2 && ref[:2] == "#/" {
		r := RemoveTrailingSlash(ref[2:])

		return strings.Split(r, "/")
	}

	return nil
}

// GetPathParamName - removing parentheses {}
func GetPathParamName(param string) string {
	if strings.HasPrefix(param, "{") && strings.HasSuffix(param, "}") {
		return param[1 : len(param)-1]
	}

	return ""
}
