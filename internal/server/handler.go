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
	handlers, err := s.Handlers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if h, ok := handlers[r.Method+" "+r.URL.Path]; ok {
		w.WriteHeader(h.statusCode)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, h.response)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func (s *Server) Handlers() (map[string]Handler, error) {
	handlers := make(map[string]Handler)
	for path, v := range s.openapi.Paths {
		for code, v := range v.Get.Responses {
			statusCode, err := strconv.Atoi(code)
			if err != nil {
				return nil, err
			}
			for _, example := range v.Content {
				handlers[http.MethodGet+" "+path] = Handler{
					method:     http.MethodGet,
					path:       path,
					statusCode: statusCode,
					response:   example.Example,
				}
			}
		}

	}

	return handlers, nil
}
