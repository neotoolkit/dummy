package server

import (
	"fmt"
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
		fmt.Fprint(w, h.response)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (s *Server) Handlers() error {
	s.handlers = make(map[string]Handler)

	for path, v := range s.openapi.Paths {
		for code, resp := range v.Get.Responses {
			statusCode, err := strconv.Atoi(code)
			if err != nil {
				return err
			}

			s.handlers[http.MethodGet+" "+path] = Handler{
				method:     http.MethodGet,
				path:       path,
				statusCode: statusCode,
				response:   resp.Content["application/json"].Example,
			}
		}
	}

	return nil
}
