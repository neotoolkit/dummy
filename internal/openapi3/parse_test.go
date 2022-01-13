package openapi3_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/api"
	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestParse_YAML(t *testing.T) {
	expected := api.API{
		Operations: []api.Operation{
			{
				Method: "POST",
				Path:   "/users",
				Body: map[string]api.FieldType{
					"id": {
						Required: true,
						Type:     "string",
					},
					"firstName": {
						Required: true,
						Type:     "string",
					},
					"lastName": {
						Required: true,
						Type:     "string",
					},
				},
				Responses: []api.Response{
					{
						StatusCode: 201,
						MediaType:  "application/json",
						Schema: api.ObjectSchema{
							Properties: map[string]api.Schema{
								"id":        api.StringSchema{Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a"},
								"firstName": api.StringSchema{Example: "Larry"},
								"lastName":  api.StringSchema{Example: "Page"},
							},
							Example: map[string]interface{}{},
						},
						Examples: map[string]interface{}{},
					},
				},
			},
			{
				Method: "GET",
				Path:   "/users",
				Responses: []api.Response{
					{
						StatusCode: 200,
						MediaType:  "application/json",
						Schema: api.ArraySchema{
							Type: api.ObjectSchema{
								Properties: map[string]api.Schema{
									"id":        api.StringSchema{Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a"},
									"firstName": api.StringSchema{Example: "Larry"},
									"lastName":  api.StringSchema{Example: "Page"},
								},
								Example: map[string]interface{}{},
							},
							Example: []interface{}{},
						},
						Example: []map[string]interface{}{
							{
								"id":        "e1afccea-5168-4735-84d4-cb96f6fb5d25",
								"firstName": "Elon",
								"lastName":  "Musk",
							},
							{
								"id":        "472063cc-4c83-11ec-81d3-0242ac130003",
								"firstName": "Sergey",
								"lastName":  "Brin",
							},
						},
						Examples: map[string]interface{}{},
					},
				},
			},
			{
				Method: "GET",
				Path:   "/users/{userId}",
				Responses: []api.Response{
					{
						StatusCode: 200,
						MediaType:  "application/json",
						Schema: api.ObjectSchema{
							Properties: map[string]api.Schema{
								"id":        api.StringSchema{Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a"},
								"firstName": api.StringSchema{Example: "Larry"},
								"lastName":  api.StringSchema{Example: "Page"},
							},
							Example: map[string]interface{}{},
						},
						Examples: map[string]interface{}{},
					},
				},
			},
		},
	}

	openapi, err := openapi3.Parse("testdata/openapi3.yml")

	require.NoError(t, err)
	require.Equalf(t, testable(expected), testable(openapi), `parsed schema from "testdata/openapi3.yml"`)
}

func testable(api api.API) api.API {
	sort.Slice(api.Operations, func(i, j int) bool {
		a, b := api.Operations[i], api.Operations[j]

		if a.Method > b.Method {
			return false
		}

		if a.Method < b.Method {
			return true
		}

		return a.Path < b.Path
	})

	return api
}
