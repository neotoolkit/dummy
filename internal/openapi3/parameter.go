package openapi3

type Parameter struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	In   string `json:"in,omitempty" yaml:"in,omitempty"`
}

type Parameters []Parameter
