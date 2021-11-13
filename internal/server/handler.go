package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	method     string
	path       string
	statusCode int
	response   interface{}
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if h, ok := s.handlers[r.Method+" "+r.URL.Path]; ok {
		w.WriteHeader(h.statusCode)
		bytes, _ := json.Marshal(h.response)
		w.Write(bytes)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (s *Server) Handlers() error {
	s.handlers = make(map[string]Handler)

	for path, method := range s.openapi.Paths {
		for code, resp := range method.Get.Responses {
			statusCode, err := strconv.Atoi(code)
			if err != nil {
				return err
			}

			s.handlers[http.MethodGet+" "+path] = Handler{
				method:     http.MethodGet,
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
