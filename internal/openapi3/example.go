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

func ExampleToResponse(data interface{}) interface{} {
	switch data := data.(type) {
	case map[interface{}]interface{}:
		return parseExample(data)
	case []interface{}:
		res := make([]map[string]interface{}, len(data))
		for k, v := range data {
			res[k] = parseExample(v.(map[interface{}]interface{}))
		}

		return res
	case string:
		return data
	}

	return nil
}

func parseExample(example map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{}, len(example))
	for k, v := range example {
		res[k.(string)] = v
	}

	return res
}
