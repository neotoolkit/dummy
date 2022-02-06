package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	gModel "github.com/neotoolkit/dummy/internal/graphql/model"
	"github.com/neotoolkit/dummy/internal/graphql/parser"
	"github.com/neotoolkit/dummy/internal/model"
)

type API struct {
	root gModel.ObjectType
	ts   *gModel.TypeSystem
}

func (a API) FindResponse(params model.FindResponseParams) (model.Response, error) {
	type GraphqlRequest struct {
		Query string `json:"query"`
	}
	req := GraphqlRequest{}
	err := json.NewDecoder(params.Body).Decode(&req)
	if err != nil {
		return nil, errors.WithMessage(err, "read body")
	}

	p := parser.NewQueryParser(parser.NewLexer(req.Query))
	query := p.Parse()

	respObject := make(map[string]interface{})
	for _, selectionField := range query.SelectionSet {
		rootObjectField, ok := a.root.FieldByName(selectionField.Name)
		if !ok {
			return errorResponse{fmt.Sprintf("uknown selectionField %q", selectionField.Name)}, nil
		}

		objectType, ok := a.ts.ObjectTypeByName(rootObjectField.Type)
		if !ok {
			return errorResponse{fmt.Sprintf("uknown object type %q", rootObjectField.Type)}, nil
		}

		fields := objectType.Fields
		if len(selectionField.SelectionSet) > 0 {
			fields = selectFields(fields, selectionField.SelectionSet)
		}

		putFieldsToResponse(respObject, rootObjectField.Name, fields)
	}

	return response{respObject: respObject}, nil
}

func selectFields(fields []gModel.Field, selectionSet gModel.SelectionSet) []gModel.Field {
	selected := make([]gModel.Field, 0, len(selectionSet))
	for _, field := range fields {
		if in(field, selectionSet) {
			selected = append(selected, field)
		}
	}
	return selected
}

func in(field gModel.Field, selectionSet gModel.SelectionSet) bool {
	for _, f := range selectionSet {
		if f.Name == field.Name {
			return true
		}
	}
	return false
}

func putFieldsToResponse(object map[string]interface{}, key string, fields []gModel.Field) {
	nestedObject := make(map[string]interface{})
	for _, field := range fields {
		nestedObject[field.Name] = fmt.Sprintf("<%s>", field.Type)
	}
	object[key] = nestedObject
}

type response struct {
	respObject map[string]interface{}
}

func (r response) StatusCode() int {
	return http.StatusOK
}

func (r response) ExampleValue(_ string) interface{} {
	return map[string]interface{}{
		"data": r.respObject,
	}
}

type errorResponse struct {
	message string
}

func (e errorResponse) StatusCode() int {
	return http.StatusOK
}

func (e errorResponse) ExampleValue(_ string) interface{} {
	return map[string]interface{}{
		"errors": map[string]interface{}{
			"message": e.message,
		},
	}
}
