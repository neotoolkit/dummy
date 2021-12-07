package server

import (
	"encoding/json"
	"fmt"
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

	for p, handlers := range h.Handlers {
		for i := 0; i < len(handlers); i++ {
			if handlers[i].Method == http.MethodGet {
				if LastPathSegmentIsParam(p) && handlers[i].Response == nil {
					handler, found := h.GetByPathAndMethod(ParentPath(p), http.MethodGet)
					if found {
						response := handler.Response.([]map[string]interface{})
						for i := 0; i < len(response); i++ {
							id, found := response[i]["id"]
							if found {
								path := ParentPath(p) + "/" + id.(string)
								h.Handlers[path] = append(h.Handlers[path], h.set(path, http.MethodGet, url.Values{}, map[string]string{}, 200, response[i]))
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func (h Handlers) GetByPathAndMethod(path, method string) (*Handler, bool) {
	handlers, found := h.Handlers[path]
	if found {
		for i := 0; i < len(handlers); i++ {
			if handlers[i].Method == method {
				return &handlers[i], true
			}
		}
	}

	return nil, false
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
			res = append(res, h.set(path, method, queryParam, map[string]string{}, statusCode, content.ResponseByExamplesKey(examplesKeys[0])))

			for i := 0; i < len(examplesKeys); i++ {
				res = append(res, h.set(path, method, queryParam, map[string]string{"X-Example": examplesKeys[i]}, statusCode, content.ResponseByExamplesKey(examplesKeys[i])))
			}
		} else {
			res = append(res, h.set(path, method, queryParam, map[string]string{}, statusCode, content.ResponseByExample()))
		}
	}

	return res, nil
}

func (h Handlers) set(path, method string, queryParam url.Values, header map[string]string, statusCode int, response interface{}) Handler {
	return Handler{
		Path:       path,
		Method:     method,
		QueryParam: queryParam,
		Header:     header,
		StatusCode: statusCode,
		Response:   response,
	}
}

func (h Handlers) Get(path, method string, queryParam url.Values, exampleHeader string, body io.ReadCloser) (Handler, bool) {
	for p, handlers := range h.Handlers {
		if PathByParamDetect(path, p) {
			for i := 0; i < len(handlers); i++ {
				if handlers[i].Method == method {
					header, ok := handlers[i].Header["X-Example"]
					if ok && header == exampleHeader {
						return handlers[i], true
					}
				}
			}

			for i := 0; i < len(handlers); i++ {
				if handlers[i].Method == method {
					if LastPathSegmentIsParam(p) && handlers[i].Response == nil {
						handler, found := h.GetByPathAndMethod(ParentPath(p), method)
						if found {
							response := handler.Response.([]map[string]interface{})
							for i := 0; i < len(response); i++ {
								if response[i]["id"] == GetLastPathSegment(path) {
									h.Handlers[path] = append(h.Handlers[path], h.set(path, method, url.Values{}, map[string]string{}, 200, response[i]))

									return h.Handlers[path][0], true
								}
							}

							return Handler{}, false
						}
					}

					if method == http.MethodPost {
						handler, found := h.GetByPathAndMethod(path, http.MethodGet)
						if found {
							data, ok := handler.Response.([]map[string]interface{})
							if ok {
								var res map[string]interface{}

								err := json.NewDecoder(body).Decode(&res)
								if err != nil {
									fmt.Println(err)
								}

								data = append(data, res)

								handler.Response = data

								return Handler{
									StatusCode: http.StatusCreated,
									Response:   res,
								}, true
							}
						}
					}

					if method == http.MethodPut || method == http.MethodPatch {
						handler, found := h.GetByPathAndMethod(path, http.MethodGet)
						if found {
							if _, ok := handler.Response.(map[string]interface{}); ok {
								var res map[string]interface{}

								err := json.NewDecoder(body).Decode(&res)
								if err != nil {
									fmt.Println(err)
								}

								handler.Response = res

								return Handler{
									StatusCode: http.StatusOK,
									Response:   res,
								}, true
							}
						}
					}

					limit, offset, found, err := pagination(queryParam)
					if err != nil {
						fmt.Println(err)
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
