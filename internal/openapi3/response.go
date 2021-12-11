package openapi3

type Response struct {
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
}

type Responses map[string]*Response
