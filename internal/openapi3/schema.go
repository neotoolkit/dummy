package openapi3

import "fmt"

type Schema struct {
	Properties Schemas `json:"properties,omitempty" yaml:"properties,omitempty"`
	Type       string  `json:"type,omitempty" yaml:"type,omitempty"`
	Format     string  `json:"format,omitempty" yaml:"format,omitempty"`
	Default    any     `json:"default,omitempty" yaml:"default,omitempty"`
	Example    any     `json:"example,omitempty" yaml:"example,omitempty"`
	Faker      string  `json:"x-faker,omitempty" yaml:"x-faker,omitempty"`

	Items *Schema `json:"items,omitempty" yaml:"items,omitempty"`

	Reference string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Schemas map[string]*Schema

type SchemaContext interface {
	LookupByReference(ref string) (Schema, error)
}

func (s Schema) ResponseByExample(schemaContext SchemaContext) (any, error) {
	if s.Reference != "" {
		schema, err := schemaContext.LookupByReference(s.Reference)
		if err != nil {
			return nil, fmt.Errorf("lookup: %w", err)
		}

		return schema.ResponseByExample(schemaContext)
	}

	if s.Example != nil {
		return ExampleToResponse(s.Example), nil
	}

	return s.propertiesExamples(schemaContext)
}

func (s Schema) propertiesExamples(schemaContext SchemaContext) (any, error) {
	if s.Items != nil {
		resp, err := s.Items.ResponseByExample(schemaContext)

		if err != nil {
			return nil, fmt.Errorf("response from items: %w", err)
		}

		var res []any
		res = append(res, resp)

		return res, nil
	}

	res := make(map[string]any, len(s.Properties))

	for key, prop := range s.Properties {
		propResp, err := prop.ResponseByExample(schemaContext)
		if err != nil {
			return nil, fmt.Errorf("response for property %q: %w", key, err)
		}

		res[key] = propResp
	}

	return res, nil
}
