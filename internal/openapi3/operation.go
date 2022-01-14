package openapi3

// Operation -.
type Operation struct {
	Parameters  Parameters  `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody RequestBody `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   Responses   `json:"responses" yaml:"responses"`
}
