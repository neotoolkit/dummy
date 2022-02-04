package api_test

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/neotoolkit/openapi"
	"github.com/stretchr/testify/require"

	"github.com/neotoolkit/dummy/internal/api"
)

func TestSchemaTypeError(t *testing.T) {
	got := &api.SchemaTypeError{
		SchemaType: "",
	}

	require.Equal(t, got.Error(), "unknown type ")
}

func TestArrayExampleError(t *testing.T) {
	got := &api.ArrayExampleError{
		Data: "",
	}

	require.Equal(t, got.Error(), "unpredicted type for example string")
}

func TestParseArrayExample(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want []interface{}
		err  error
	}{
		{
			name: "nil data",
			data: nil,
			want: []interface{}{},
			err:  nil,
		},
		{
			name: "array",
			data: []interface{}{
				map[string]interface{}{
					"key": "value",
				},
			},
			want: []interface{}{
				map[string]interface{}{
					"key": "value",
				},
			},
			err: nil,
		},
		{
			name: "not array",
			data: "string",
			want: nil,
			err:  &api.ArrayExampleError{Data: "string"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := api.ParseArrayExample(tc.data)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}

			require.Equal(t, tc.want, res)
		})
	}
}

func TestObjectExampleError(t *testing.T) {
	got := &api.ObjectExampleError{
		Data: "",
	}

	require.Equal(t, got.Error(), "unpredicted type for example string")
}

func TestParseObjectExample(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want map[string]interface{}
		err  error
	}{
		{
			name: "nil data",
			data: nil,
			want: map[string]interface{}{},
			err:  nil,
		},
		{
			name: "object",
			data: map[string]interface{}{},
			want: map[string]interface{}{},
			err:  nil,
		},
		{
			name: "not object",
			data: "string",
			want: nil,
			err:  &api.ObjectExampleError{Data: "string"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := api.ParseObjectExample(tc.data)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}

			require.Equal(t, tc.want, res)
		})
	}
}

func TestRemoveTrailingSlash(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "",
			path: "",
			want: "",
		},
		{
			name: "",
			path: "/",
			want: "",
		},
		{
			name: "",
			path: "path/",
			want: "path",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := api.RemoveTrailingSlash(tc.path)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	tests := []struct {
		name    string
		builder api.Builder
		want    api.API
		err     error
	}{
		{
			name:    "",
			builder: api.Builder{},
			want:    api.API{},
			err:     nil,
		},
		{
			name: "GET",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Get: &openapi.Operation{},
						},
					},
				},
			},
			want: api.API{
				Operations: []api.Operation{
					{
						Method:    "GET",
						Path:      "test",
						Body:      map[string]api.FieldType(nil),
						Responses: []api.Response(nil),
					},
				},
			},
			err: nil,
		},
		{
			name: "Wrong status code in GET",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Get: &openapi.Operation{
								Responses: map[string]*openapi.Response{
									"Wrong status code": nil,
								},
							},
						},
					},
				},
			},
			want: api.API{},
			err: &strconv.NumError{
				Func: "Atoi",
				Num:  "Wrong status code",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			name: "POST",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Post: &openapi.Operation{},
						},
					},
				},
			},
			want: api.API{
				Operations: []api.Operation{
					{
						Method:    "POST",
						Path:      "test",
						Body:      map[string]api.FieldType(nil),
						Responses: []api.Response(nil),
					},
				},
			},
			err: nil,
		},
		{
			name: "Wrong schema reference in POST",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Post: &openapi.Operation{
								RequestBody: openapi.RequestBody{
									Content: map[string]*openapi.MediaType{
										"application/json": {
											Schema: openapi.Schema{
												Ref: "wrong schema reference",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: api.API{},
			err:  fmt.Errorf("resolve reference: %w", &openapi.SchemaError{Ref: "wrong schema reference"}),
		},
		{
			name: "PUT",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Put: &openapi.Operation{},
						},
					},
				},
			},
			want: api.API{
				Operations: []api.Operation{
					{
						Method:    "PUT",
						Path:      "test",
						Body:      map[string]api.FieldType(nil),
						Responses: []api.Response(nil),
					},
				},
			},
			err: nil,
		},
		{
			name: "Wrong schema reference in PUT",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Put: &openapi.Operation{
								RequestBody: openapi.RequestBody{
									Content: map[string]*openapi.MediaType{
										"application/json": {
											Schema: openapi.Schema{
												Ref: "wrong schema reference",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: api.API{},
			err:  fmt.Errorf("resolve reference: %w", &openapi.SchemaError{Ref: "wrong schema reference"}),
		},
		{
			name: "PATCH",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Patch: &openapi.Operation{},
						},
					},
				},
			},
			want: api.API{
				Operations: []api.Operation{
					{
						Method:    "PATCH",
						Path:      "test",
						Body:      map[string]api.FieldType(nil),
						Responses: []api.Response(nil),
					},
				},
			},
			err: nil,
		},
		{
			name: "Wrong schema reference in PATCH",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Patch: &openapi.Operation{
								RequestBody: openapi.RequestBody{
									Content: map[string]*openapi.MediaType{
										"application/json": {
											Schema: openapi.Schema{
												Ref: "wrong schema reference",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: api.API{},
			err:  fmt.Errorf("resolve reference: %w", &openapi.SchemaError{Ref: "wrong schema reference"}),
		},
		{
			name: "DELETE",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Delete: &openapi.Operation{},
						},
					},
				},
			},
			want: api.API{
				Operations: []api.Operation{
					{
						Method:    "DELETE",
						Path:      "test",
						Body:      map[string]api.FieldType(nil),
						Responses: []api.Response(nil),
					},
				},
			},
			err: nil,
		},
		{
			name: "Wrong status code in DELETE",
			builder: api.Builder{
				OpenAPI: openapi.OpenAPI{
					Paths: map[string]*openapi.Path{
						"test": {
							Delete: &openapi.Operation{
								Responses: map[string]*openapi.Response{
									"Wrong status code": nil,
								},
							},
						},
					},
				},
			},
			want: api.API{},
			err: &strconv.NumError{
				Func: "Atoi",
				Num:  "Wrong status code",
				Err:  strconv.ErrSyntax,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.builder

			res, err := b.Build()
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}

			require.Equal(t, tc.want, res)
		})
	}
}

func TestBuilder_Add(t *testing.T) {
	tests := []struct {
		name      string
		builder   api.Builder
		path      string
		method    string
		operation *openapi.Operation
		err       error
	}{
		{
			name:      "",
			builder:   api.Builder{},
			path:      "",
			method:    "",
			operation: nil,
			err:       nil,
		},
		{
			name:    "",
			builder: api.Builder{},
			path:    "",
			method:  "",
			operation: &openapi.Operation{
				RequestBody: openapi.RequestBody{
					Content: map[string]*openapi.MediaType{
						"application/json": {
							Schema: openapi.Schema{
								Ref: "wrong schema reference",
							},
						},
					},
				},
			},
			err: fmt.Errorf("resolve reference: %w", &openapi.SchemaError{Ref: "wrong schema reference"}),
		},
		{
			name:      "",
			builder:   api.Builder{},
			path:      "",
			method:    "",
			operation: &openapi.Operation{},
			err:       nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.builder

			err := b.Add(tc.path, tc.method, tc.operation)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}
		})
	}
}

func TestBuilder_Set(t *testing.T) {
	tests := []struct {
		name      string
		builder   api.Builder
		path      string
		method    string
		operation *openapi.Operation
		want      api.Operation
		err       error
	}{
		{
			name:      "",
			builder:   api.Builder{},
			path:      "",
			method:    "",
			operation: nil,
			want:      api.Operation{},
			err:       nil,
		},
		{
			name:    "",
			builder: api.Builder{},
			path:    "",
			method:  "",
			operation: &openapi.Operation{
				RequestBody: openapi.RequestBody{
					Content: map[string]*openapi.MediaType{
						"application/json": {
							Schema: openapi.Schema{
								Ref: "wrong schema reference",
							},
						},
					},
				},
			},
			want: api.Operation{},
			err:  fmt.Errorf("resolve reference: %w", &openapi.SchemaError{Ref: "wrong schema reference"}),
		},
		{
			name:    "",
			builder: api.Builder{},
			path:    "",
			method:  "",
			operation: &openapi.Operation{
				RequestBody: openapi.RequestBody{
					Content: map[string]*openapi.MediaType{
						"application/json": {
							Schema: openapi.Schema{
								Required: []string{"field"},
								Properties: map[string]*openapi.Schema{
									"prop": {},
								},
							},
						},
					},
				},
			},
			want: api.Operation{
				Body: map[string]api.FieldType{
					"field": {
						Required: true,
						Type:     "",
					},
					"prop": {
						Required: false,
						Type:     "",
					},
				},
			},
			err: nil,
		},
		{
			name:    "",
			builder: api.Builder{},
			path:    "",
			method:  "",
			operation: &openapi.Operation{
				Responses: openapi.Responses{
					"200": {},
				},
			},
			want: api.Operation{
				Responses: []api.Response{
					{
						StatusCode: http.StatusOK,
					},
				},
			},
			err: nil,
		},
		{
			name:    "",
			builder: api.Builder{},
			path:    "",
			method:  "",
			operation: &openapi.Operation{
				Responses: openapi.Responses{
					"200": {
						Content: map[string]*openapi.MediaType{
							"application/json": {
								Example: "example",
							},
						},
					},
				},
			},
			want: api.Operation{},
			err:  &api.SchemaTypeError{SchemaType: ""},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := tc.builder

			got, err := b.Set(tc.path, tc.method, tc.operation)
			if err != nil {
				require.EqualError(t, err, tc.err.Error())
			}

			require.Equal(t, tc.want, got)
		})
	}
}
