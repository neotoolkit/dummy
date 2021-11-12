package openapi3

type Path struct {
	Get *Operation `json:"get,omitempty" yaml:"get,omitempty"`
}

type Paths map[string]*Path
