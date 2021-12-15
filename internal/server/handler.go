package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-dummy/dummy/internal/logger"
	"github.com/go-dummy/dummy/internal/openapi3"
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
	OpenAPI  *openapi3.OpenAPI
	Handlers map[string][]Handler
	Logger   *logger.Logger
}

// NewHandlers returns a new instance of Handlers
func NewHandlers(openapi *openapi3.OpenAPI, l *logger.Logger) Handlers {
	return Handlers{
		OpenAPI:  openapi,
		Handlers: make(map[string][]Handler),
		Logger:   l,
	}
}

// Init -
func (h Handlers) Init() error {
	for path, method := range h.OpenAPI.Paths {
		if err := h.Add(path, http.MethodGet, method.Get); err != nil {
			return err
		}

		if err := h.Add(path, http.MethodPost, method.Post); err != nil {
			return err
		}

		if err := h.Add(path, http.MethodPut, method.Put); err != nil {
			return err
		}

		if err := h.Add(path, http.MethodPatch, method.Patch); err != nil {
			return err
		}

		if err := h.Add(path, http.MethodDelete, method.Delete); err != nil {
			return err
		}
	}

	return nil
}

// Add -.
func (h Handlers) Add(path, method string, o *openapi3.Operation) error {
	if o != nil {
		p := RemoveTrailingSlash(path)

		handlers, err := h.Set(p, method, o)
		if err != nil {
			return err
		}

		h.Handlers[p] = append(h.Handlers[p], handlers...)
	}

	return nil
}

// Set -.
func (h Handlers) Set(path, method string, o *openapi3.Operation) ([]Handler, error) {
	var res []Handler

	queryParam := make(url.Values)

	for i := 0; i < len(o.Parameters); i++ {
		if o.Parameters[i].In == "query" {
			queryParam.Add(o.Parameters[i].Name, "")
		}
	}

	for code, resp := range o.Responses {
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			return nil, err
		}

		content, found := resp.Content["application/json"]
		if found {
			examplesKeys := content.Examples.GetKeys()

			switch {
			case len(examplesKeys) > 0:
				res = append(res, h.set(path, method, queryParam, http.Header{}, statusCode, content.ResponseByExamplesKey(examplesKeys[0])))

				for i := 0; i < len(examplesKeys); i++ {
					res = append(res, h.set(path, method, queryParam, http.Header{"X-Example": []string{examplesKeys[i]}}, statusCode, content.ResponseByExamplesKey(examplesKeys[i])))
				}
			case content.Example != nil:
				res = append(res, h.set(path, method, queryParam, http.Header{}, statusCode, content.ResponseByExample()))
			default:
				schemaResp, err := content.Schema.ResponseByExample(h.OpenAPI)
				if err != nil {
					return nil, fmt.Errorf("response from schema: %w", err)
				}

				res = append(res, h.set(path, method, queryParam, http.Header{}, statusCode, schemaResp))
			}
		} else {
			res = append(res, h.set(path, method, queryParam, http.Header{}, statusCode, nil))
		}
	}

	return res, nil
}

func (h Handlers) set(path, method string, queryParam url.Values, header http.Header, statusCode int, response any) Handler {
	return Handler{
		Path:       path,
		Method:     method,
		QueryParam: queryParam,
		Header:     header,
		StatusCode: statusCode,
		Response:   response,
	}
}

// Get -.
func (h Handlers) Get(path, method string, queryParam url.Values, header http.Header, body io.ReadCloser) (Handler, bool) {
	for p, handlers := range h.Handlers {
		if PathByParamDetect(path, p) {
			for i := 0; i < len(handlers); i++ {
				if handlers[i].Method == method && reflect.DeepEqual(handlers[i].Header.Values("X-Example"), header.Values("X-Example")) {
					return handlers[i], true
				}
			}
		}
	}

	return Handler{}, false
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
