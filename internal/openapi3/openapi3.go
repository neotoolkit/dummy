package openapi3

// OpenAPI Object
// See specification https://swagger.io/specification/#openapi-object
type OpenAPI struct {
	Info       Info       `json:"info" yaml:"info"`
	Paths      Paths      `json:"paths" yaml:"paths"`
	Components Components `json:"components,omitempty" yaml:"components,omitempty"`
}
