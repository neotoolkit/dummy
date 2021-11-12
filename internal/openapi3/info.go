package openapi3

type Info struct {
	Title          string  `json:"title" yaml:"title"`
	Description    string  `json:"description" yaml:"description"`
	TermsOfService string  `json:"termsOfService" yaml:"termsOfService"`
	Contact        Contact `json:"contact" yaml:"contact"`
	License        License `json:"license" yaml:"license"`
	Version        string  `json:"version" yaml:"version"`
}
