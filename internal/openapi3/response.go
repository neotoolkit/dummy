package openapi3

// Response -.
type Response struct {
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Content     Content `json:"content,omitempty" yaml:"content,omitempty"`
}

// Responses -.
type Responses map[string]*Response
