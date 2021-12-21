package apischema

import (
	"fmt"
	"strings"
)

type FindResponseParams struct {
	Path      string
	Method    string
	MediaType string
}

func (a API) FindResponse(params FindResponseParams) (Response, error) {
	operation, found := a.findOperation(params)
	if !found {
		return Response{}, fmt.Errorf("not specified operation %q", params.Method+" "+params.Path)
	}

	response, found := operation.findResponse(params)
	if !found {
		return operation.Responses[0], nil
	}

	return response, nil
}

func (a API) findOperation(params FindResponseParams) (Operation, bool) {
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

func (o Operation) findResponse(params FindResponseParams) (Response, bool) {
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
