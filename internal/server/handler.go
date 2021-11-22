package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-dummy/dummy/internal/openapi3"
)

type Handler struct {
	Path       string
	Method     string
	Header     map[string]string
	StatusCode int
	Response   interface{}
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if h, ok := s.GetHandler(r.Method, r.URL.Path, r.Header.Get("example")); ok {
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

	for code, resp := range o.Responses {
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			return nil, err
		}

		content := resp.Content["application/json"]

		examplesKeys := getExamplesKeys(content.Examples)

		if len(examplesKeys) > 0 {
			res = append(res, handler(path, method, map[string]string{}, statusCode, response(content, examplesKeys[0])))

			for i := 0; i < len(examplesKeys); i++ {
				res = append(res, handler(path, method, map[string]string{"example": examplesKeys[i]}, statusCode, response(content, examplesKeys[i])))
			}
		} else {
			res = append(res, handler(path, method, map[string]string{}, statusCode, response(content)))
		}
	}

	return res, nil
}

func handler(path, method string, header map[string]string, statusCode int, response interface{}) Handler {
	return Handler{
		Path:       path,
		Method:     method,
		Header:     header,
		StatusCode: statusCode,
		Response:   response,
	}
}

func (s *Server) GetHandler(method, path, exampleHeader string) (h Handler, found bool) {
	for mask, handlers := range s.Handlers {
		if pathMaskDetect(path, mask) {
			for i := 0; i < len(handlers); i++ {
				if handlers[i].Method == method {
					for header, v := range handlers[i].Header {
						if header == "example" && v == exampleHeader {
							h = handlers[i]
							found = true

							return
						}
					}

					if lastParamIsMask(mask) && handlers[i].Response == nil {
						for _, v := range s.Handlers[parentPath(mask)] {
							if v.Method == method {
								data := v.Response.([]map[string]interface{})
								for i := 0; i < len(data); i++ {
									if data[i]["id"] == getLastParam(path) {
										s.Handlers[path] = append(s.Handlers[path], handler(path, method, map[string]string{}, 200, data[i]))

										h = s.Handlers[path][0]
										found = true

										return
									}
								}
							}
						}
					}

					h = handlers[i]
					found = true
				}
			}
		}
	}

	return
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

func getLastParam(path string) string {
	p := strings.Split(path, "/")

	return p[len(p)-1]
}
