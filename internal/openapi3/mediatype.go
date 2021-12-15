package openapi3

type MediaType struct {
	Example  any      `json:"example,omitempty" yaml:"example,omitempty"`
	Examples Examples `json:"examples,omitempty" yaml:"examples,omitempty"`
}

func (mt MediaType) ResponseByExample() any {
	return mt.example(mt.Example)
}

func (mt MediaType) ResponseByExamplesKey(key string) any {
	return mt.examples(key)
}

func (mt MediaType) example(i any) any {
	switch data := i.(type) {
	case map[any]any:
		return parseExample(data)
	case []any:
		res := make([]map[string]any, len(data))
		for k, v := range data {
			res[k] = parseExample(v.(map[any]any))
		}

		return res
	}

	return nil
}

func (mt MediaType) examples(key string) any {
	return mt.example(mt.Examples[key].Value)
}

func parseExample(example map[any]any) map[string]any {
	res := make(map[string]any, len(example))
	for k, v := range example {
		res[k.(string)] = v
	}

	return res
}
