package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/neotoolkit/faker"
	"github.com/neotoolkit/openapi"

	"github.com/neotoolkit/dummy/internal/pkg/pathfmt"
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
	Data any
}

func (e *ArrayExampleError) Error() string {
	return fmt.Sprintf("unpredicted type for example %T", e.Data)
}

func ParseArrayExample(data any) ([]any, error) {
	if nil == data {
		return []any{}, nil
	}

	d, ok := data.([]any)
	if ok {
		res := make([]any, len(d))
		for k, v := range d {
			res[k] = v.(map[string]any)
		}

		return res, nil
	}

	return nil, &ArrayExampleError{Data: data}
}

// ObjectExampleError -.
type ObjectExampleError struct {
	Data any
}

// Error -.
func (e *ObjectExampleError) Error() string {
	return fmt.Sprintf("unpredicted type for example %T", e.Data)
}

func ParseObjectExample(data any) (map[string]any, error) {
	if nil == data {
		return map[string]any{}, nil
	}

	d, ok := data.(map[string]any)
	if ok {
		return d, nil
	}

	return nil, &ObjectExampleError{Data: data}
}

type Builder struct {
	OpenAPI    openapi.OpenAPI
	Operations []Operation
	Faker      faker.Faker
}

// Build -.
func (b *Builder) Build() (API, error) {
	for path, method := range b.OpenAPI.Paths {
		if err := b.Add(path, http.MethodGet, method.Get); err != nil {
			return API{}, err
		}

		if err := b.Add(path, http.MethodPost, method.Post); err != nil {
			return API{}, err
		}

		if err := b.Add(path, http.MethodPut, method.Put); err != nil {
			return API{}, err
		}

		if err := b.Add(path, http.MethodPatch, method.Patch); err != nil {
			return API{}, err
		}

		if err := b.Add(path, http.MethodDelete, method.Delete); err != nil {
			return API{}, err
		}
	}

	return API{Operations: b.Operations}, nil
}

// Add -.
func (b *Builder) Add(path, method string, o *openapi.Operation) error {
	if o != nil {
		p := pathfmt.RemoveTrailingSlash(path)

		operation, err := b.Set(p, method, o)
		if err != nil {
			return err
		}

		b.Operations = append(b.Operations, operation)
	}

	return nil
}

// Set -.
func (b *Builder) Set(path, method string, o *openapi.Operation) (Operation, error) {
	operation := Operation{
		Method: method,
		Path:   path,
	}

	if nil == o {
		return operation, nil
	}

	body, ok := o.RequestBody.Content["application/json"]
	if ok || len(o.RequestBody.Ref) > 0 {
		var s openapi.Schema

		switch {
		case len(o.RequestBody.Ref) > 0:
			requestBody, err := b.OpenAPI.LookupRequestBodyByReference(o.RequestBody.Ref)
			if err != nil {
				return Operation{}, fmt.Errorf("resolve reference: %w", err)
			}

			body, ok := requestBody.Content["application/json"]
			if ok {
				s = body.Schema
			}
		case body.Schema.IsRef():
			schema, err := b.OpenAPI.LookupSchemaByReference(body.Schema.Ref)
			if err != nil {
				return Operation{}, fmt.Errorf("resolve reference: %w", err)
			}

			s = schema
		default:
			s = body.Schema
		}

		operation.Body = make(map[string]FieldType, len(s.Properties))

		for _, v := range s.Required {
			operation.Body[v] = FieldType{Required: true}
		}

		for k, v := range s.Properties {
			operation.Body[k] = FieldType{
				Required: operation.Body[k].Required,
				Type:     v.Type,
			}
		}
	}

	for code, resp := range o.Responses {
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			return Operation{}, err
		}

		if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {
			continue
		}

		content, ok := resp.Content["application/json"]
		if !ok {
			operation.Responses = append(operation.Responses, Response{
				StatusCode: statusCode,
			})

			continue
		}

		example := openapi.ExampleToResponse(content.Example)

		examples := make(map[string]any, len(content.Examples)+1)

		if len(content.Examples) > 0 {
			for key, e := range content.Examples {
				examples[key] = openapi.ExampleToResponse(e.Value)
			}

			examples[""] = openapi.ExampleToResponse(content.Examples[content.Examples.GetKeys()[0]].Value)
		}

		schema, err := b.convertSchema(content.Schema)
		if err != nil {
			return Operation{}, err
		}

		operation.Responses = append(operation.Responses, Response{
			StatusCode: statusCode,
			MediaType:  "application/json",
			Schema:     schema,
			Example:    example,
			Examples:   examples,
		})
	}

	return operation, nil
}

func (b *Builder) convertSchema(s openapi.Schema) (Schema, error) {
	if s.IsRef() {
		schema, err := b.OpenAPI.LookupSchemaByReference(s.Ref)
		if err != nil {
			return nil, fmt.Errorf("resolve reference: %w", err)
		}

		s = schema
	}

	if s.Faker != "" {
		return FakerSchema{Example: b.Faker.ByName(s.Faker)}, nil
	}

	switch s.Type {
	case "boolean":
		val, _ := s.Example.(bool)
		return BooleanSchema{Example: val}, nil
	case "integer":
		val, _ := s.Example.(int64)
		return IntSchema{Example: val}, nil
	case "number":
		val, _ := s.Example.(float64)
		return FloatSchema{Example: val}, nil
	case "string":
		val, _ := s.Example.(string)
		return StringSchema{Example: val}, nil
	case "array":
		if nil == s.Items {
			return nil, ErrEmptyItems
		}

		itemsSchema, err := b.convertSchema(*s.Items)
		if err != nil {
			return nil, err
		}

		arrExample, err := ParseArrayExample(s.Example)
		if err != nil {
			return nil, err
		}

		return ArraySchema{
			Type:    itemsSchema,
			Example: arrExample,
		}, nil
	case "object":
		obj := ObjectSchema{Properties: make(map[string]Schema, len(s.Properties))}

		for key, prop := range s.Properties {
			propSchema, err := b.convertSchema(*prop)
			if err != nil {
				return nil, err
			}

			obj.Properties[key] = propSchema
		}

		objExample, err := ParseObjectExample(s.Example)
		if err != nil {
			return nil, err
		}

		obj.Example = objExample

		return obj, nil
	default:
		return nil, &SchemaTypeError{SchemaType: s.Type}
	}
}
