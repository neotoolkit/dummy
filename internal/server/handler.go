package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-dummy/dummy/internal/openapi3"
)

type Handler struct {
	method     string
	path       string
	statusCode int
	response   interface{}
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if h, ok := s.Handlers[r.Method+" "+r.URL.Path]; ok {
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

		if statusCode >= http.StatusOK || statusCode <= http.StatusNoContent {
			h[method+" "+path] = Handler{
				method:     method,
				path:       path,
				statusCode: statusCode,
				response:   example(resp.Content["application/json"].Example),
			}
		}
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
