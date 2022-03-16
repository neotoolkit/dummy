package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neotoolkit/dummy/internal/api"
)

func TestSchema_ExampleValue(t *testing.T) {
	tests := []struct {
		name   string
		schema api.Schema
		want   any
	}{
		{
			name:   "boolean: default",
			schema: api.BooleanSchema{},
			want:   false,
		},
		{
			name:   "boolean: with example",
			schema: api.BooleanSchema{Example: true},
			want:   true,
		},
		{
			name:   "int: default",
			schema: api.IntSchema{},
			want:   int64(0),
		},
		{
			name:   "int: with example",
			schema: api.IntSchema{Example: 42},
			want:   int64(42),
		},
		{
			name:   "float: default",
			schema: api.FloatSchema{},
			want:   0.0,
		},
		{
			name:   "float: with example",
			schema: api.FloatSchema{Example: 4.2},
			want:   4.2,
		},
		{
			name:   "string: default",
			schema: api.StringSchema{},
			want:   "",
		},
		{
			name:   "string: with example",
			schema: api.StringSchema{Example: "John"},
			want:   "John",
		},
		{
			name:   "array: default",
			schema: api.ArraySchema{Type: api.StringSchema{}},
			want:   []any{""},
		},
		{
			name:   "array: with int example",
			schema: api.ArraySchema{Example: []any{4, 2}},
			want:   []any{4, 2},
		},
		{
			name:   "array: with string example",
			schema: api.ArraySchema{Example: []any{"4", "2"}},
			want:   []any{"4", "2"},
		},
		{
			name:   "object: default",
			schema: api.ObjectSchema{},
			want:   map[string]any{},
		},
		{
			name:   "object: with example",
			schema: api.ObjectSchema{Example: map[string]any{"a": "4", "b": "2"}},
			want:   map[string]any{"a": "4", "b": "2"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.schema.ExampleValue()

			require.Equal(t, tc.want, got)
		})
	}
}

func TestResponse_ExampleValue(t *testing.T) {
	tests := []struct {
		name     string
		response api.Response
		key      string
		want     any
	}{
		{
			name: "boolean examples",
			response: api.Response{
				Schema:   api.BooleanSchema{},
				Examples: map[string]any{"first": false, "second": true},
			},
			key:  "first",
			want: false,
		},
		{
			name: "int examples",
			response: api.Response{
				Schema:   api.IntSchema{},
				Examples: map[string]any{"first": int64(4), "second": int64(2)},
			},
			key:  "first",
			want: int64(4),
		},
		{
			name: "float examples",
			response: api.Response{
				Schema:   api.FloatSchema{},
				Examples: map[string]any{"first": 4.0, "second": 2.0},
			},
			key:  "first",
			want: 4.0,
		},
		{
			name: "string examples",
			response: api.Response{
				Schema:   api.StringSchema{},
				Examples: map[string]any{"first": "abc", "second": "xyz"},
			},
			key:  "first",
			want: "abc",
		},
		{
			name: "array examples",
			response: api.Response{
				Schema:   api.ArraySchema{},
				Examples: map[string]any{"first": []int64{4, 2}, "second": []int64{0, 0}},
			},
			key:  "first",
			want: []int64{4, 2},
		},
		{
			name: "object examples",
			response: api.Response{
				Schema:   api.ObjectSchema{},
				Examples: map[string]any{"first": map[string]any{"first": "abc"}, "second": map[string]any{"first": "xyz"}},
			},
			key:  "first",
			want: map[string]interface{}{"first": "abc"},
		},
		{
			name: "use schema example",
			response: api.Response{
				Examples: map[string]any{"first": map[string]any{"first": "abc"}, "second": map[string]any{"first": "xyz"}},
				Schema: api.ObjectSchema{
					Example: map[string]any{"first": "schema variant"},
				},
			},
			key:  "third",
			want: map[string]any{"first": "schema variant"},
		},
		{
			name: "nil schema",
			response: api.Response{
				Schema:   nil,
				Examples: map[string]any{"first": false, "second": true},
			},
			key:  "first",
			want: nil,
		},
		{
			name: "example",
			response: api.Response{
				Schema:  api.ObjectSchema{},
				Example: map[string]any{"first": "abc", "second": "def"},
			},
			key:  "",
			want: map[string]any{"first": "abc", "second": "def"},
		},
		{
			name: "object properties",
			response: api.Response{
				Schema: api.ObjectSchema{
					Properties: map[string]api.Schema{
						"key": api.ObjectSchema{},
					},
				},
			},
			key:  "",
			want: map[string]any{"key": map[string]any{}},
		},
		{
			name: "faker schema",
			response: api.Response{
				Schema: api.FakerSchema{},
			},
			key:  "",
			want: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.response.ExampleValue(tc.key)

			require.Equal(t, tc.want, got)
		})
	}
}
