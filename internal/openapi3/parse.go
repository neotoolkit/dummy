package openapi3

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/go-dummy/dummy/internal/api"
	"github.com/go-dummy/dummy/internal/faker"
	"github.com/go-dummy/dummy/internal/read"
)

// Parse -.
func Parse(path string) (api.API, error) {
	file, err := read.Read(path)
	if err != nil {
		return api.API{}, err
	}

	var openapi OpenAPI

	if err := yaml.Unmarshal(file, &openapi); err != nil {
		return api.API{}, err
	}

	f := faker.NewFaker()

	b := &builder{openapi: openapi, faker: f}

	return b.Build()
}

// SchemaTypeError -.
type SchemaTypeError struct {
	schemaType string
}

func (e *SchemaTypeError) Error() string {
	return "unknown type " + e.schemaType
}

// ErrEmptyItems -.
var ErrEmptyItems = errors.New("empty items in array")

// ArrayExampleError -.
type ArrayExampleError struct {
	data interface{}
}

func (e *ArrayExampleError) Error() string {
	return fmt.Sprintf("unpredicted type for example %T", e.data)
}

func parseArrayExample(data interface{}) ([]interface{}, error) {
	if nil == data {
		return []interface{}{}, nil
	}

	d, ok := data.([]interface{})
	if ok {
		res := make([]interface{}, len(d))
		for k, v := range d {
			res[k] = v.(map[string]interface{})
		}

		return res, nil
	}

	return nil, &ArrayExampleError{data: data}
}

// ObjectExampleError -.
type ObjectExampleError struct {
	data interface{}
}

// Error -.
func (e *ObjectExampleError) Error() string {
	return fmt.Sprintf("unpredicted type for example %T", e.data)
}

func parseObjectExample(data interface{}) (map[string]interface{}, error) {
	if nil == data {
		return map[string]interface{}{}, nil
	}

	d, ok := data.(map[string]interface{})
	if ok {
		return d, nil
	}

	return nil, &ObjectExampleError{data: data}
}

// RemoveTrailingSlash returns path without trailing slash
func RemoveTrailingSlash(path string) string {
	if len(path) > 0 && path[len(path)-1] == '/' {
		return path[0 : len(path)-1]
	}

	return path
}
