package openapi3

type Path struct {
	Post *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Get  *Operation `json:"get,omitempty" yaml:"get,omitempty"`
}

type Paths map[string]*Path
