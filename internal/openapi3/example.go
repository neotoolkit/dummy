package openapi3

// Example -.
type Example struct {
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

// Examples -.
type Examples map[string]Example

// GetKeys -.
func (e Examples) GetKeys() []string {
	keys := make([]string, len(e))
	i := 0

	for k := range e {
		keys[i] = k
		i++
	}

	return keys
}

// ExampleToResponse -.
func ExampleToResponse(data interface{}) interface{} {
	switch d := data.(type) {
	case map[string]interface{}:
		return d
	case []interface{}:
		res := make([]map[string]interface{}, len(d))
		for k, v := range d {
			res[k] = v.(map[string]interface{})
		}

		return res
	case string:
		return d
	}

	return nil
}
