package openapi3

type Schema struct {
	Properties Schemas     `json:"properties,omitempty" yaml:"properties,omitempty"`
	Type       string      `json:"type,omitempty" yaml:"type,omitempty"`
	Format     string      `json:"format,omitempty" yaml:"format,omitempty"`
	Example    interface{} `json:"example,omitempty" yaml:"example,omitempty"`
}

type Schemas map[string]*Schema
