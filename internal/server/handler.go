package server

import (
	"encoding/json"
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
	Header     map[string]string
	StatusCode int
	Response   interface{}
}

type Handlers struct {
	OpenAPI  openapi3.OpenAPI
	Handlers map[string][]Handler
}

func (h Handlers) Set() error {
	for path, method := range h.OpenAPI.Paths {
		if method.Get != nil {
			handlers, err := handlers(path, http.MethodGet, method.Get)
			if err != nil {
				return err
			}

			h.Handlers[path] = append(h.Handlers[path], handlers...)
		}

		if method.Post != nil {
			handlers, err := handlers(path, http.MethodPost, method.Post)
			if err != nil {
				return err
			}

			h.Handlers[path] = append(h.Handlers[path], handlers...)
		}

		if method.Put != nil {
			handlers, err := handlers(path, http.MethodPut, method.Put)
			if err != nil {
				return err
			}

			h.Handlers[path] = append(h.Handlers[path], handlers...)
		}

		if method.Patch != nil {
			handlers, err := handlers(path, http.MethodPatch, method.Patch)
			if err != nil {
				return err
			}

			h.Handlers[path] = append(h.Handlers[path], handlers...)
		}
	}

	return nil
}

func handlers(path, method string, o *openapi3.Operation) ([]Handler, error) {
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
			res = append(res, handler(path, method, queryParam, map[string]string{}, statusCode, content.ResponseByExamplesKey(examplesKeys[0])))

			for i := 0; i < len(examplesKeys); i++ {
				res = append(res, handler(path, method, queryParam, map[string]string{"example": examplesKeys[i]}, statusCode, content.ResponseByExamplesKey(examplesKeys[i])))
			}
		} else {
			res = append(res, handler(path, method, queryParam, map[string]string{}, statusCode, content.ResponseByExample()))
		}
	}

	return res, nil
}

func handler(path, method string, queryParam url.Values, header map[string]string, statusCode int, response interface{}) Handler {
	return Handler{
		Path:       path,
		Method:     method,
		QueryParam: queryParam,
		Header:     header,
		StatusCode: statusCode,
		Response:   response,
	}
}

func (s *Server) GetHandler(method, path string, queryParam url.Values, exampleHeader string, body io.ReadCloser) (Handler, bool) {
	for mask, handlers := range s.Handlers.Handlers {
		if PathByParamDetect(path, mask) {
			for i := 0; i < len(handlers); i++ {
				if handlers[i].Method == method {
					header, ok := handlers[i].Header["example"]
					if ok && header == exampleHeader {
						return handlers[i], true
					}
				}
			}

			for i := 0; i < len(handlers); i++ {
				if handlers[i].Method == method {
					if LastPathSegmentIsParam(mask) && handlers[i].Response == nil {
						h, found := s.getHandlerByPathAndMethod(ParentPath(mask), method)
						if found {
							data := h.Response.([]map[string]interface{})
							for i := 0; i < len(data); i++ {
								if data[i]["id"] == GetLastPathSegment(path) {
									s.Handlers.Handlers[path] = append(s.Handlers.Handlers[path], handler(path, method, url.Values{}, map[string]string{}, 200, data[i]))

									return s.Handlers.Handlers[path][0], true
								}
							}

							return Handler{}, false
						}
					}

					if method == http.MethodPost {
						h, found := s.getHandlerByPathAndMethod(path, http.MethodGet)
						if found {
							data, ok := h.Response.([]map[string]interface{})
							if ok {
								var res map[string]interface{}

								err := json.NewDecoder(body).Decode(&res)
								if err != nil {
									s.Logger.Log().Err(err)
								}

								data = append(data, res)

								h.Response = data

								return Handler{
									StatusCode: http.StatusCreated,
									Response:   res,
								}, true
							}
						}
					}

					if method == http.MethodPut || method == http.MethodPatch {
						h, found := s.getHandlerByPathAndMethod(path, http.MethodGet)
						if found {
							if _, ok := h.Response.(map[string]interface{}); ok {
								var res map[string]interface{}

								err := json.NewDecoder(body).Decode(&res)
								if err != nil {
									s.Logger.Log().Err(err)
								}

								h.Response = res

								return Handler{
									StatusCode: http.StatusOK,
									Response:   res,
								}, true
							}
						}
					}

					limit, offset, found, err := pagination(queryParam)
					if err != nil {
						s.Logger.Log().Err(err)
					}

					if found {
						data := handlers[i].Response.([]map[string]interface{})
						if offset > len(data) {
							return Handler{}, true
						}

						size := len(data) - offset
						if size > limit {
							size = limit
						}

						resp := make([]map[string]interface{}, size)
						n := 0

						for i := offset; n < size; i++ {
							resp[n] = data[i]
							n++
						}

						return Handler{
							Response:   resp,
							StatusCode: http.StatusOK,
						}, true
					}

					return handlers[i], true
				}
			}
		}
	}

	return Handler{}, false
}

func (s *Server) getHandlerByPathAndMethod(path, method string) (*Handler, bool) {
	h, found := s.Handlers.Handlers[path]
	if found {
		for i := 0; i < len(h); i++ {
			if h[i].Method == method {
				return &h[i], true
			}
		}
	}

	return &Handler{}, false
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

func pagination(queryParam url.Values) (limit, offset int, found bool, err error) {
	l, ok := queryParam["limit"]
	if !ok {
		return
	}

	if ok {
		limit, err = strconv.Atoi(l[0])
	}

	o, ok := queryParam["offset"]
	if ok {
		offset, err = strconv.Atoi(o[0])
	}

	found = true

	return
}

func RemoveTrailingSlash(path string) string {
	if len(path) > 0 && path[len(path)-1] == '/' {
		return path[0 : len(path)-1]
	}

	return path
}

// RemoveFragment - clear url from reference in path
func RemoveFragment(path string) string {
	return strings.Split(path, "#")[0]
}
