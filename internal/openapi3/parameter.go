package openapi3

// Parameter -.
type Parameter struct {
	Name     string  `json:"name,omitempty" yaml:"name,omitempty"`
	In       string  `json:"in,omitempty" yaml:"in,omitempty"`
	Required bool    `json:"required,omitempty" yaml:"required,omitempty"`
	Schema   *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

// Parameters -.
type Parameters []Parameter
