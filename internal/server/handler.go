package server

import (
	"encoding/json"
	"net/http"
	"reflect"
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

			val := reflect.ValueOf(resp.Content["application/json"].Example)

			var res interface{}

			switch val.Kind() {
			case reflect.Map:
				m := make(map[string]interface{}, len(val.MapKeys()))
				for _, e := range val.MapKeys() {
					m[e.Elem().String()] = val.MapIndex(e).Elem().String()
				}
				res = m
			case reflect.Slice:
				arr := make([]interface{}, 0)
				for i := 0; i < val.Len(); i++ {
					m := make(map[string]interface{})
					for _, e := range val.Index(i).Elem().MapKeys() {
						m[e.Elem().String()] = val.Index(i).Elem().MapIndex(e).Elem().String()
					}
					arr = append(arr, m)
				}
				res = arr
			}

			s.handlers[http.MethodGet+" "+path] = Handler{
				method:     http.MethodGet,
				path:       path,
				statusCode: statusCode,
				response:   res,
			}
		}
	}

	return nil
}
