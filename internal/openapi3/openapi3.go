package openapi3

// The root document object of the OpenAPI document.
type OpenAPI struct {
	OpenAPI string `json:"openapi" yaml:"openapi"`
	Info    *Info  `json:"info" yaml:"info"`
}
