package openapi3

type Example struct {
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

type Examples map[string]Example

func (e Examples) GetKeys() []string {
	keys := make([]string, len(e))
	i := 0

	for k := range e {
		keys[i] = k
		i++
	}

	return keys
}
