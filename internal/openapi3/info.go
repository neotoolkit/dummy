package openapi3

// The object provides metadata about the API. The metadata MAY be used by the clients if needed, and MAY be presented in editing or documentation generation tools for convenience.
type Info struct {
	// The title of the API.
	Title string `json:"title" yaml:"title"`
	// A short description of the API.
	Description string `json:"description" yaml:"description"`
	// A URL to the Terms of Service for the API.
	TermsOfService string `json:"termsOfService" yaml:"termsOfService"`
	// The contact information for the exposed API.
	Contact Contact `json:"contact" yaml:"contact"`
	// The license information for the exposed API.
	License License `json:"license" yaml:"license"`
	// The version of the OpenAPI document.
	Version string `json:"version" yaml:"version"`
}

// Contact information for the exposed API.
type Contact struct {
	// The identifying name of the contact person/organization
	Name string `json:"name" yaml:"name"`
	// The URL pointing to the contact information
	URL string `json:"url" yaml:"url"`
	// The email address of the contact person/organization.
	Email string `json:"email" yaml:"email"`
}

// License information for the exposed API.
type License struct {
	// The license name used for the API.
	Name string `json:"name" yaml:"name"`
	// A URL to the license used for the API.
	URL string `json:"url" yaml:"url"`
}
