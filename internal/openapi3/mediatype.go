package openapi3

type MediaType struct {
	Example  interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	Examples Examples    `json:"examples,omitempty" yaml:"examples,omitempty"`
}
