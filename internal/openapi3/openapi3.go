package openapi3

// The root document object of the OpenAPI document.
type OpenAPI struct {
	// This string MUST be the semantic version number of the OpenAPI Specification version that the OpenAPI document uses. The openapi field SHOULD be used by tooling specifications and clients to interpret the OpenAPI document. This is not related to the API info.version string.
	OpenAPI string `json:"openapi" yaml:"openapi"`
	// Provides metadata about the API. The metadata MAY be used by tooling as required.
	Info *Info `json:"info" yaml:"info"`
}
