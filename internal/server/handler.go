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

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if h, ok := s.GetHandler(r.Method, r.URL.Path, r.URL.Query(), r.Header.Get("x-example"), r.Body); ok {
		w.WriteHeader(h.StatusCode)
		bytes, _ := json.Marshal(h.Response)
		_, _ = w.Write(bytes)

		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (s *Server) SetHandlers() error {
	for path, method := range s.OpenAPI.Paths {
		if method.Get != nil {
			handlers, err := handlers(path, http.MethodGet, method.Get)
			if err != nil {
				return err
			}

			s.Handlers[path] = append(s.Handlers[path], handlers...)
		}

		if method.Post != nil {
			handlers, err := handlers(path, http.MethodPost, method.Post)
			if err != nil {
				return err
			}

			s.Handlers[path] = append(s.Handlers[path], handlers...)
		}
	}

	return nil
}

func handlers(path, method string, o *openapi3.Operation) ([]Handler, error) {
	var res []Handler

	queryParam := make(url.Values)

	for i := 0; i < len(o.Parameters); i++ {
		if o.Parameters[i].In == "query" {
			queryParam[o.Parameters[i].Name] = append(queryParam[o.Parameters[i].Name], "")
		}
	}

	for code, resp := range o.Responses {
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			return nil, err
		}

		content := resp.Content["application/json"]

		examplesKeys := getExamplesKeys(content.Examples)

		if len(examplesKeys) > 0 {
			res = append(res, handler(path, method, queryParam, map[string]string{}, statusCode, response(content, examplesKeys[0])))

			for i := 0; i < len(examplesKeys); i++ {
				res = append(res, handler(path, method, queryParam, map[string]string{"example": examplesKeys[i]}, statusCode, response(content, examplesKeys[i])))
			}
		} else {
			res = append(res, handler(path, method, queryParam, map[string]string{}, statusCode, response(content)))
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
	for mask, handlers := range s.Handlers {
		if pathMaskDetect(path, mask) {
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
					if lastParamIsMask(mask) && handlers[i].Response == nil {
						for _, v := range s.Handlers[parentPath(mask)] {
							if v.Method == method {
								data := v.Response.([]map[string]interface{})
								for i := 0; i < len(data); i++ {
									if data[i]["id"] == GetLastPathSegment(path) {
										s.Handlers[path] = append(s.Handlers[path], handler(path, method, url.Values{}, map[string]string{}, 200, data[i]))

										return s.Handlers[path][0], true
									}
								}
							}
						}
					}

					if method == http.MethodPost {
						for i := 0; i < len(s.Handlers[path]); i++ {
							if s.Handlers[path][i].Method == http.MethodGet {
								data, ok := s.Handlers[path][i].Response.([]map[string]interface{})
								if ok {
									var res map[string]interface{}

									err := json.NewDecoder(body).Decode(&res)
									if err != nil {
										s.Logger.Log().Err(err)
									}

									data = append(data, res)

									s.Handlers[path][i].Response = data

									return Handler{
										StatusCode: s.Handlers[path][i].StatusCode,
										Response:   res,
									}, true
								}
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

func response(mt *openapi3.MediaType, key ...string) interface{} {
	if mt.Example != nil {
		return example(mt.Example)
	}

	if len(mt.Examples) > 0 && len(key) > 0 {
		return examples(mt.Examples, key[0])
	}

	return nil
}

func example(i interface{}) interface{} {
	switch data := i.(type) {
	case map[interface{}]interface{}:
		return parseExample(data)
	case []interface{}:
		res := make([]map[string]interface{}, len(data))
		for k, v := range data {
			res[k] = parseExample(v.(map[interface{}]interface{}))
		}

		return res
	}

	return nil
}

func parseExample(example map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{}, len(example))
	for k, v := range example {
		res[k.(string)] = v
	}

	return res
}

func examples(e openapi3.Examples, key string) interface{} {
	return example(e[key].Value)
}

func getExamplesKeys(e map[string]openapi3.Example) []string {
	keys := make([]string, len(e))
	i := 0

	for k := range e {
		keys[i] = k
		i++
	}

	return keys
}

func pathMaskDetect(path, mask string) bool {
	p := strings.Split(path, "/")
	m := strings.Split(mask, "/")

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

func parentPath(path string) string {
	p := strings.Split(path, "/")

	return strings.Join(p[0:len(p)-1], "/")
}

func lastParamIsMask(path string) bool {
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
