package openapi3

type Schema struct {
	Properties Schemas     `json:"properties,omitempty" yaml:"properties,omitempty"`
	Type       string      `json:"type,omitempty" yaml:"type,omitempty"`
	Format     string      `json:"format,omitempty" yaml:"format,omitempty"`
	Default    interface{} `json:"default,omitempty" yaml:"default,omitempty"`
	Example    interface{} `json:"example,omitempty" yaml:"example,omitempty"`
}

type Schemas map[string]*Schema
