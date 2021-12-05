package openapi3

type MediaType struct {
	Example  interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	Examples Examples    `json:"examples,omitempty" yaml:"examples,omitempty"`
}

func (mt MediaType) Response(key ...string) interface{} {
	if mt.Example != nil {
		return mt.example(mt.Example)
	}

	if len(mt.Examples) > 0 && len(key) > 0 {
		return mt.examples(key[0])
	}

	return nil
}

func (mt MediaType) example(i interface{}) interface{} {
	switch data := i.(type) {
	case map[interface{}]interface{}:
		return parseExample(data)
	case []interface{}:
		res := make([]map[string]interface{}, len(data))
		for k, v := range data {
			res[k] = parseExample(v.(map[interface{}]interface{}))
		}

		return res
	}

	return nil
}

func (mt MediaType) examples(key string) interface{} {
	return mt.example(mt.Examples[key].Value)
}

func parseExample(example map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{}, len(example))
	for k, v := range example {
		res[k.(string)] = v
	}

	return res
}
