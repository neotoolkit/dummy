package openapi3

// OpenAPI v3 document
type OpenAPI struct {
	OpenAPI string `json:"openapi" yaml:"openapi"`
	Info    *Info  `json:"info" yaml:"info"`
}
