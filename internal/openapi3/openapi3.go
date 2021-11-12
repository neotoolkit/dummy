package openapi3

type OpenAPI struct {
	OpenAPI string `json:"openapi" yaml:"openapi"`
	Info    *Info  `json:"info" yaml:"info"`
	Tags    Tags   `json:"tags,omitempty" yaml:"tags,omitempty"`
}
