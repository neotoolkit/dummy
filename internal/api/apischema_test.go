package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-dummy/dummy/internal/api"
)

func TestSchema_ExampleValue(t *testing.T) {
	tests := []struct {
		name   string
		schema api.Schema
		want   interface{}
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
			want:   []interface{}{""},
		},
		{
			name:   "array: with int example",
			schema: api.ArraySchema{Example: []interface{}{4, 2}},
			want:   []interface{}{4, 2},
		},
		{
			name:   "array: with string example",
			schema: api.ArraySchema{Example: []interface{}{"4", "2"}},
			want:   []interface{}{"4", "2"},
		},
		{
			name:   "object: default",
			schema: api.ObjectSchema{},
			want:   map[string]interface{}{},
		},
		{
			name:   "object: with example",
			schema: api.ObjectSchema{Example: map[string]interface{}{"a": "4", "b": "2"}},
			want:   map[string]interface{}{"a": "4", "b": "2"},
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
		want     interface{}
	}{
		{
			name:     "boolean examples",
			response: api.Response{Schema: api.BooleanSchema{}, Examples: map[string]interface{}{"first": false, "second": true}},
			key:      "first",
			want:     false,
		},
		{
			name:     "int examples",
			response: api.Response{Schema: api.IntSchema{}, Examples: map[string]interface{}{"first": int64(4), "second": int64(2)}},
			key:      "first",
			want:     int64(4),
		},
		{
			name:     "float examples",
			response: api.Response{Schema: api.FloatSchema{}, Examples: map[string]interface{}{"first": 4.0, "second": 2.0}},
			key:      "first",
			want:     4.0,
		},
		{
			name:     "string examples",
			response: api.Response{Schema: api.StringSchema{}, Examples: map[string]interface{}{"first": "abc", "second": "xyz"}},
			key:      "first",
			want:     "abc",
		},
		{
			name:     "array examples",
			response: api.Response{Schema: api.ArraySchema{}, Examples: map[string]interface{}{"first": []int64{4, 2}, "second": []int64{0, 0}}},
			key:      "first",
			want:     []int64{4, 2},
		},
		{
			name:     "object examples",
			response: api.Response{Schema: api.ObjectSchema{}, Examples: map[string]interface{}{"first": map[string]interface{}{"first": "abc"}, "second": map[string]interface{}{"first": "xyz"}}},
			key:      "first",
			want:     map[string]interface{}{"first": "abc"},
		},
		{
			name: "use schema example",
			response: api.Response{Examples: map[string]interface{}{"first": map[string]interface{}{"first": "abc"}, "second": map[string]interface{}{"first": "xyz"}},
				Schema: api.ObjectSchema{
					Example: map[string]interface{}{"first": "schema variant"},
				}},
			key:  "third",
			want: map[string]interface{}{"first": "schema variant"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.response.ExampleValue(tc.key)

			require.Equal(t, tc.want, got)
		})
	}
}
