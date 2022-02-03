package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/neotoolkit/dummy/internal/model"
)

// FindResponseError -.
type FindResponseError struct {
	Method string
	Path   string
}

// Error -.
func (e *FindResponseError) Error() string {
	return "not specified operation: " + e.Method + " " + e.Path
}

// ErrEmptyRequireField -.
var ErrEmptyRequireField = errors.New("empty require field")

// FindResponse -.
func (a API) FindResponse(params model.FindResponseParams) (model.Response, error) {
	operation, ok := a.findOperation(params)
	if !ok {
		return nil, &FindResponseError{
			Method: params.Method,
			Path:   params.Path,
		}
	}

	switch params.Method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		var body map[string]interface{}

		err := json.NewDecoder(params.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		for k, v := range operation.Body {
			_, ok := body[k]
			if !ok && v.Required {
				return nil, ErrEmptyRequireField
			}
		}
	}

	response, ok := operation.findResponse(params)
	if !ok {
		return operation.Responses[0], nil
	}

	return response, nil
}

func (a API) findOperation(params model.FindResponseParams) (Operation, bool) {
	for _, op := range a.Operations {
		if !PathByParamDetect(params.Path, op.Path) {
			continue
		}

		if params.Method != op.Method {
			continue
		}

		return op, true
	}

	return Operation{}, false
}

func (o Operation) findResponse(params model.FindResponseParams) (Response, bool) {
	for _, r := range o.Responses {
		if r.MediaType != params.MediaType {
			continue
		}

		return r, true
	}

	return Response{}, false
}

// PathByParamDetect returns result of
func PathByParamDetect(path, param string) bool {
	splitPath := strings.Split(path, "/")
	splitParam := strings.Split(param, "/")

	if len(splitPath) != len(splitParam) {
		return false
	}

	for i := 0; i < len(splitPath); i++ {
		if strings.HasPrefix(splitParam[i], "{") && strings.HasSuffix(splitParam[i], "}") {
			continue
		}

		if splitPath[i] != splitParam[i] {
			return false
		}
	}

	return true
}
