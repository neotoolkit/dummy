package openapi3_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/apischema"
	"github.com/go-dummy/dummy/internal/openapi3"
)

func TestParse_Yaml(t *testing.T) {
	expected := apischema.API{
		Operations: []apischema.Operation{
			{
				Method: "POST",
				Path:   "/users",
				Responses: []apischema.Response{
					{
						StatusCode: 201,
						MediaType:  "application/json",
						Schema: apischema.ObjectSchema{
							Properties: map[string]apischema.Schema{
								"id":        apischema.StringSchema{Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a"},
								"firstName": apischema.StringSchema{Example: "Larry"},
								"lastName":  apischema.StringSchema{Example: "Page"},
							},
							Example: map[string]any{},
						},
						Examples: map[string]any{},
					},
				},
			},
			{
				Method: "GET",
				Path:   "/users",
				Responses: []apischema.Response{
					{
						StatusCode: 200,
						MediaType:  "application/json",
						Schema: apischema.ArraySchema{
							Type: apischema.ObjectSchema{
								Properties: map[string]apischema.Schema{
									"id":        apischema.StringSchema{Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a"},
									"firstName": apischema.StringSchema{Example: "Larry"},
									"lastName":  apischema.StringSchema{Example: "Page"},
								},
								Example: map[string]any{},
							},
							Example: []any{},
						},
						Example: []map[string]any{
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
						Examples: map[string]any{},
					},
				},
			},
			{
				Method: "GET",
				Path:   "/users/{userId}",
				Responses: []apischema.Response{
					{
						StatusCode: 200,
						MediaType:  "application/json",
						Schema: apischema.ObjectSchema{
							Properties: map[string]apischema.Schema{
								"id":        apischema.StringSchema{Example: "380ed0b7-eb21-4ad4-acd0-efa90cf69c6a"},
								"firstName": apischema.StringSchema{Example: "Larry"},
								"lastName":  apischema.StringSchema{Example: "Page"},
							},
							Example: map[string]any{},
						},
						Examples: map[string]any{},
					},
				},
			},
		},
	}

	openapi, err := openapi3.Parse("testdata/openapi3.yml")

	require.NoError(t, err)
	require.Equalf(t, testable(expected), testable(openapi), `parsed schema from "testdata/openapi3.yml"`)
}

func testable(api apischema.API) apischema.API {
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
