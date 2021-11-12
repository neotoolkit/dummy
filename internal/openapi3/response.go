package openapi3

type Response struct {
	Content Content `json:"content,omitempty" yaml:"content,omitempty"`
}

type Responses map[string]*Response
