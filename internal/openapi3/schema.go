package openapi3

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
