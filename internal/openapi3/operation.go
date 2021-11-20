package openapi3

type Operation struct {
	Parameters Parameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses  Responses  `json:"responses" yaml:"responses"`
}
