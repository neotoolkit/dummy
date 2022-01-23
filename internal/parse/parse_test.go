package parse_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/api"
	"github.com/go-dummy/dummy/internal/parse"
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

	openapi, err := parse.Parse("testdata/openapi3.yml")

	require.NoError(t, err)
	require.Equalf(t, testable(t, expected), testable(t, openapi), `parsed schema from "testdata/openapi3.yml"`)
}

func testable(t *testing.T, api api.API) api.API {
	t.Helper()

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

func TestGetSpecType(t *testing.T) {
	tests := []struct {
		name string
		path string
		want parse.SpecType
		err  error
	}{
		{
			name: "",
			path: "",
			want: parse.Unknown,
			err:  parse.ErrEmptySpecTypePath,
		},
		{
			name: "",
			path: "./testdata/openapi3",
			want: parse.Unknown,
			err: &parse.SpecFileError{
				Path: "./testdata/openapi3",
			},
		},
		{
			name: "",
			path: "./testdata/openapi3.yml",
			want: parse.OpenAPI,
			err:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parse.GetSpecType(tc.path)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}
			require.Equal(t, tc.want, got)
		})
	}
}
