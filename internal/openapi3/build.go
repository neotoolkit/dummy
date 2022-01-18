package openapi3

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-dummy/dummy/internal/api"
	"github.com/go-dummy/dummy/internal/faker"
)

type builder struct {
	openapi    OpenAPI
	operations []api.Operation
	faker      faker.Faker
}

// Build -.
func (b *builder) Build() (api.API, error) {
	for path, method := range b.openapi.Paths {
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

	return api.API{Operations: b.operations}, nil
}

// Add -.
func (b *builder) Add(path, method string, o *Operation) error {
	if o != nil {
		p := RemoveTrailingSlash(path)

		operation, err := b.Set(p, method, o)
		if err != nil {
			return err
		}

		b.operations = append(b.operations, operation)
	}

	return nil
}

// Set -.
func (b *builder) Set(path, method string, o *Operation) (api.Operation, error) {
	operation := api.Operation{
		Method: method,
		Path:   path,
	}

	body, ok := o.RequestBody.Content["application/json"]
	if ok {
		var s Schema

		if body.Schema.Reference != "" {
			schema, err := b.openapi.LookupByReference(body.Schema.Reference)
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

func (b *builder) convertSchema(s Schema) (api.Schema, error) {
	if s.Reference != "" {
		schema, err := b.openapi.LookupByReference(s.Reference)
		if err != nil {
			return nil, fmt.Errorf("resolve reference: %w", err)
		}

		s = schema
	}

	if s.Faker != "" {
		return api.FakerSchema{Example: b.faker.ByName(s.Faker)}, nil
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
