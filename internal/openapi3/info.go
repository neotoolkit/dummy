package openapi3

type Info struct {
	Title          string  `json:"title" yaml:"title"`
	Description    string  `json:"description" yaml:"description"`
	TermsOfService string  `json:"termsOfService" yaml:"termsOfService"`
	Contact        Contact `json:"contact" yaml:"contact"`
	License        License `json:"license" yaml:"license"`
	Version        string  `json:"version" yaml:"version"`
}

type Contact struct {
	Name  string `json:"name" yaml:"name"`
	URL   string `json:"url" yaml:"url"`
	Email string `json:"email" yaml:"email"`
}

type License struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}
