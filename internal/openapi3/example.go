package openapi3

type Example struct {
	Value any `json:"value,omitempty" yaml:"value,omitempty"`
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

func ExampleToResponse(data any) any {
	switch data := data.(type) {
	case map[any]any:
		return parseExample(data)
	case []any:
		res := make([]map[string]any, len(data))
		for k, v := range data {
			res[k] = parseExample(v.(map[any]any))
		}

		return res
	case string:
		return data
	}

	return nil
}

func parseExample(example map[any]any) map[string]any {
	res := make(map[string]any, len(example))
	for k, v := range example {
		res[k.(string)] = v
	}

	return res
}
