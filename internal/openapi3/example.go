package openapi3

type Example struct {
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

type Examples map[string]Example
