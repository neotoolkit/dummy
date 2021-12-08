package server

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-dummy/dummy/internal/openapi3"
)

type Handler struct {
	Path       string
	Method     string
	QueryParam url.Values
	Header     http.Header
	StatusCode int
	Response   interface{}
}

type Handlers struct {
	OpenAPI  openapi3.OpenAPI
	Handlers map[string][]Handler
}

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
	}

	return nil
}

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

		content := resp.Content["application/json"]

		examplesKeys := content.Examples.GetExamplesKeys()

		if len(examplesKeys) > 0 {
			res = append(res, h.set(path, method, queryParam, http.Header{}, statusCode, content.ResponseByExamplesKey(examplesKeys[0])))

			for i := 0; i < len(examplesKeys); i++ {
				res = append(res, h.set(path, method, queryParam, http.Header{"X-Example": []string{examplesKeys[i]}}, statusCode, content.ResponseByExamplesKey(examplesKeys[i])))
			}
		} else {
			res = append(res, h.set(path, method, queryParam, http.Header{}, statusCode, content.ResponseByExample()))
		}
	}

	return res, nil
}

func (h Handlers) set(path, method string, queryParam url.Values, header http.Header, statusCode int, response interface{}) Handler {
	return Handler{
		Path:       path,
		Method:     method,
		QueryParam: queryParam,
		Header:     header,
		StatusCode: statusCode,
		Response:   response,
	}
}

func (h Handlers) Get(path, method string, queryParam url.Values, header http.Header, body io.ReadCloser) (Handler, bool) {
	for p, handlers := range h.Handlers {
		if PathByParamDetect(path, p) {
			for i := 0; i < len(handlers); i++ {
				if handlers[i].Method == method {
					if EqualHeadersByValues(handlers[i].Header.Values("X-Example"), header.Values("X-Example")) {
						return handlers[i], true
					}
				}
			}
		}
	}

	return Handler{}, false
}

func PathByParamDetect(path, param string) bool {
	p := strings.Split(path, "/")
	m := strings.Split(param, "/")

	if len(p) != len(m) {
		return false
	}

	for i := 0; i < len(p); i++ {
		if strings.HasPrefix(m[i], "{") && strings.HasSuffix(m[i], "}") {
			continue
		}

		if p[i] != m[i] {
			return false
		}
	}

	return true
}

func ParentPath(path string) string {
	p := strings.Split(path, "/")

	return strings.Join(p[0:len(p)-1], "/")
}

func LastPathSegmentIsParam(path string) bool {
	p := strings.Split(path, "/")

	return strings.HasPrefix(p[len(p)-1], "{") && strings.HasSuffix(p[len(p)-1], "}")
}

func GetLastPathSegment(path string) string {
	p := strings.Split(path, "/")

	return p[len(p)-1]
}

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

// EqualHeadersByValues - comparing two headers by values
func EqualHeadersByValues(h1, h2 []string) bool {
	if h1 == nil && h2 == nil {
		return true
	}

	if len(h1) != len(h2) {
		return false
	}

	exists := make(map[string]bool, len(h1))
	for i := 0; i < len(h1); i++ {
		exists[h1[i]] = true
	}

	for i := 0; i < len(h2); i++ {
		if !exists[h2[i]] {
			return false
		}
	}

	return true
}
