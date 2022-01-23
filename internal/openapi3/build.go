package openapi3

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-dummy/faker"

	"github.com/go-dummy/dummy/internal/api"
)

// SchemaTypeError -.
type SchemaTypeError struct {
	SchemaType string
}

func (e *SchemaTypeError) Error() string {
	return "unknown type " + e.SchemaType
}

// ErrEmptyItems -.
var ErrEmptyItems = errors.New("empty items in array")

// ArrayExampleError -.
type ArrayExampleError struct {
	Data interface{}
}

func (e *ArrayExampleError) Error() string {
	return fmt.Sprintf("unpredicted type for example %T", e.Data)
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

	return nil, &ArrayExampleError{Data: data}
}

// ObjectExampleError -.
type ObjectExampleError struct {
	Data interface{}
}

// Error -.
func (e *ObjectExampleError) Error() string {
	return fmt.Sprintf("unpredicted type for example %T", e.Data)
}

func parseObjectExample(data interface{}) (map[string]interface{}, error) {
	if nil == data {
		return map[string]interface{}{}, nil
	}

	d, ok := data.(map[string]interface{})
	if ok {
		return d, nil
	}

	return nil, &ObjectExampleError{Data: data}
}

// RemoveTrailingSlash returns path without trailing slash
func RemoveTrailingSlash(path string) string {
	if len(path) > 0 && path[len(path)-1] == '/' {
		return path[0 : len(path)-1]
	}

	return path
}

type Builder struct {
	OpenAPI    OpenAPI
	Operations []api.Operation
	Faker      faker.Faker
}

// Build -.
func (b *Builder) Build() (api.API, error) {
	for path, method := range b.OpenAPI.Paths {
		if err := b.Add(path, http.MethodGet, method.Get); err != nil {
			return api.API{}, err
		}

		if err := b.Add(path, http.MethodPost, method.Post); err != nil {
			return api.API{}, err
		}

		if err := b.Add(path, http.MethodPut, method.Put); err != nil {
			return api.API{}, err
		}

		if err := b.Add(path, http.MethodPatch, method.Patch); err != nil {
			return api.API{}, err
		}

		if err := b.Add(path, http.MethodDelete, method.Delete); err != nil {
			return api.API{}, err
		}
	}

	return api.API{Operations: b.Operations}, nil
}

// Add -.
func (b *Builder) Add(path, method string, o *Operation) error {
	if o != nil {
		p := RemoveTrailingSlash(path)

		operation, err := b.Set(p, method, o)
		if err != nil {
			return err
		}

		b.Operations = append(b.Operations, operation)
	}

	return nil
}

// Set -.
func (b *Builder) Set(path, method string, o *Operation) (api.Operation, error) {
	operation := api.Operation{
		Method: method,
		Path:   path,
	}

	body, ok := o.RequestBody.Content["application/json"]
	if ok {
		var s Schema

		if body.Schema.Reference != "" {
			schema, err := b.OpenAPI.LookupByReference(body.Schema.Reference)
			if err != nil {
				return api.Operation{}, fmt.Errorf("resolve reference: %w", err)
			}

			s = schema
		} else {
			s = body.Schema
		}

		operation.Body = make(map[string]api.FieldType, len(s.Properties))

		for _, v := range s.Required {
			operation.Body[v] = api.FieldType{
				Required: true,
			}
		}

		for k, v := range s.Properties {
			operation.Body[k] = api.FieldType{
				Required: operation.Body[k].Required,
				Type:     v.Type,
			}
		}
	}

	for code, resp := range o.Responses {
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			return api.Operation{}, err
		}

		content, ok := resp.Content["application/json"]
		if !ok {
			operation.Responses = append(operation.Responses, api.Response{
				StatusCode: statusCode,
			})

			continue
		}

		example := ExampleToResponse(content.Example)

		examples := make(map[string]interface{}, len(content.Examples)+1)

		if len(content.Examples) > 0 {
			for key, e := range content.Examples {
				examples[key] = ExampleToResponse(e.Value)
			}

			examples[""] = ExampleToResponse(content.Examples[content.Examples.GetKeys()[0]].Value)
		}

		schema, err := b.convertSchema(content.Schema)
		if err != nil {
			return api.Operation{}, err
		}

		operation.Responses = append(operation.Responses, api.Response{
			StatusCode: statusCode,
			MediaType:  "application/json",
			Schema:     schema,
			Example:    example,
			Examples:   examples,
		})
	}

	return operation, nil
}

func (b *Builder) convertSchema(s Schema) (api.Schema, error) {
	if s.Reference != "" {
		schema, err := b.OpenAPI.LookupByReference(s.Reference)
		if err != nil {
			return nil, fmt.Errorf("resolve reference: %w", err)
		}

		s = schema
	}

	if s.Faker != "" {
		return api.FakerSchema{Example: b.Faker.ByName(s.Faker)}, nil
	}

	switch s.Type {
	case "boolean":
		val, _ := s.Example.(bool)
		return api.BooleanSchema{Example: val}, nil
	case "integer":
		val, _ := s.Example.(int64)
		return api.IntSchema{Example: val}, nil
	case "number":
		val, _ := s.Example.(float64)
		return api.FloatSchema{Example: val}, nil
	case "string":
		val, _ := s.Example.(string)
		return api.StringSchema{Example: val}, nil
	case "array":
		if nil == s.Items {
			return nil, ErrEmptyItems
		}

		itemsSchema, err := b.convertSchema(*s.Items)
		if err != nil {
			return nil, err
		}

		arrExample, err := parseArrayExample(s.Example)
		if err != nil {
			return nil, err
		}

		return api.ArraySchema{
			Type:    itemsSchema,
			Example: arrExample,
		}, nil
	case "object":
		obj := api.ObjectSchema{Properties: make(map[string]api.Schema, len(s.Properties))}

		for key, prop := range s.Properties {
			propSchema, err := b.convertSchema(*prop)
			if err != nil {
				return nil, err
			}

			obj.Properties[key] = propSchema
		}

		objExample, err := parseObjectExample(s.Example)
		if err != nil {
			return nil, err
		}

		obj.Example = objExample

		return obj, nil
	default:
		return nil, &SchemaTypeError{SchemaType: s.Type}
	}
}
