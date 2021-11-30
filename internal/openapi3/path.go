package openapi3

type Path struct {
	Post  *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Get   *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Put   *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Patch *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
}

type Paths map[string]*Path
