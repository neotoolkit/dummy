package openapi3

// Media Type Object
// See specification https://swagger.io/specification/#media-type-object
type MediaType struct {
	Schema   Schema      `json:"schema" yaml:"schema"`
	Example  interface{} `json:"example,omitempty" yaml:"example,omitempty"`
	Examples Examples    `json:"examples,omitempty" yaml:"examples,omitempty"`
}

func (mt MediaType) ResponseByExample() interface{} {
	return ExampleToResponse(mt.Example)
}

func (mt MediaType) ResponseByExamplesKey(key string) interface{} {
	return mt.examples(key)
}

func (mt MediaType) examples(key string) interface{} {
	return ExampleToResponse(mt.Examples[key].Value)
}
