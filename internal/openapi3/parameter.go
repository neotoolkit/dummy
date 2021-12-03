package openapi3

type Parameter struct {
	Name     string  `json:"name,omitempty" yaml:"name,omitempty"`
	In       string  `json:"in,omitempty" yaml:"in,omitempty"`
	Required bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Schema   *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

type Parameters []Parameter
