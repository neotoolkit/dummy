package openapi3

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var expected = OpenAPI{
	Info: Info{
		Title:   "Users dummy API",
		Version: "0.1.0",
	},
	Paths: Paths{
		"/users": &Path{
			Post: &Operation{
				Parameters: nil,
				Responses: Responses{
					"201": &Response{
						Description: ptrStr(""),
						Content: Content{
							"application/json": &MediaType{
								Example:  nil,
								Examples: nil,
							},
						},
					},
				},
			},
			Get: &Operation{
				Parameters: nil,
				Responses: Responses{
					"200": &Response{
						Description: ptrStr(""),
						Content: Content{
							"application/json": &MediaType{
								Example: []interface{}{
									map[interface{}]interface{}{
										"id":        "e1afccea-5168-4735-84d4-cb96f6fb5d25",
										"firstName": "Elon",
										"lastName":  "Musk",
									},
									map[interface{}]interface{}{
										"id":        "472063cc-4c83-11ec-81d3-0242ac130003",
										"firstName": "Sergey",
										"lastName":  "Brin",
									},
								},
								Examples: nil,
							},
						},
					},
				},
			},
		},
		"/users/{userId}": &Path{
			Get: &Operation{
				Parameters: Parameters{
					Parameter{
						Name:     "userId",
						In:       "path",
						Required: true,
						Schema: &Schema{
							Type: "string",
						},
					},
				},
				Responses: Responses{
					"200": &Response{
						Description: ptrStr(""),
						Content: Content{
							"application/json": &MediaType{},
						},
					},
				},
			},
		},
	},
	Components: Components{
		Schemas: Schemas{
			"User": &Schema{
				Properties: Schemas{
					"id": &Schema{
						Type:   "string",
						Format: "uuid",
					},
					"firstName": &Schema{
						Type: "string",
					},
					"lastName": &Schema{
						Type: "string",
					},
				},
				Type: "object",
			},
		},
	},
}

func TestParse_Yaml(t *testing.T) {
	openapi, err := Parse("testdata/openapi3.yml")

	require.NoError(t, err)
	require.Equalf(t, expected, openapi, "parsed schema from 'testdata/openapi3.yml'")
}

func ptrStr(s string) *string {
	return &s
}
