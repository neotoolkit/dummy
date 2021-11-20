package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-dummy/dummy/internal/openapi3"
)

type Handler struct {
	statusCode int
	response   interface{}
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var key strings.Builder

	key.WriteString(r.Method + " " + r.URL.Path)

	exampleHeader := r.Header.Get("example")
	if len(exampleHeader) > 0 {
		key.WriteString("?example=" + exampleHeader)
	}

	if h, ok := s.Handlers[key.String()]; ok {
		w.WriteHeader(h.statusCode)
		bytes, _ := json.Marshal(h.response)
		_, _ = w.Write(bytes)

		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (s *Server) SetHandlers() error {
	for path, method := range s.OpenAPI.Paths {
		if err := addHandler(s.Handlers, http.MethodGet, path, method.Get); err != nil {
			return err
		}

		if err := addHandler(s.Handlers, http.MethodPost, path, method.Post); err != nil {
			return err
		}
	}

	return nil
}

func addHandler(h map[string]Handler, method, path string, o *openapi3.Operation) error {
	if o == nil {
		return nil
	}

	for code, resp := range o.Responses {
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			return err
		}

		key := method + " " + path

		if statusCode >= http.StatusOK || statusCode <= http.StatusNoContent {
			content := resp.Content["application/json"]
			keys := getExamplesKeys(content.Examples)

			if len(keys) > 0 {
				h[key] = handler(statusCode, response(content, keys[0]))

				for i := 0; i < len(keys); i++ {
					h[key+"?example="+keys[i]] = handler(statusCode, response(content, keys[i]))
				}
			} else {
				h[key] = handler(statusCode, response(content))
			}
		}
	}

	return nil
}

func handler(statusCode int, response interface{}) Handler {
	return Handler{
		statusCode: statusCode,
		response:   response,
	}
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
