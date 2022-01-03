package apischema

type API struct {
	Operations []Operation
}

type Operation struct {
	Method    string
	Path      string
	Responses []Response
}

type Response struct {
	StatusCode int
	MediaType  string
	Schema     Schema
	Example    interface{}
	Examples   map[string]interface{}
}

func (r Response) ExampleValue(key string) interface{} {
	if nil == r.Schema {
		return nil
	}

	if example, found := r.Examples[key]; found {
		return example
	}

	if r.Example != nil {
		return r.Example
	}

	return r.Schema.ExampleValue()
}

type Schema interface {
	ExampleValue() interface{}
}

type BooleanSchema struct {
	Example bool
}

func (b BooleanSchema) ExampleValue() interface{} {
	return b.Example
}

type IntSchema struct {
	Example int64
}

func (i IntSchema) ExampleValue() interface{} {
	return i.Example
}

type FloatSchema struct {
	Example float64
}

func (f FloatSchema) ExampleValue() interface{} {
	return f.Example
}

type StringSchema struct {
	Example string
}

func (s StringSchema) ExampleValue() interface{} {
	return s.Example
}

type ArraySchema struct {
	Type    Schema
	Example []interface{}
}

func (a ArraySchema) ExampleValue() interface{} {
	if len(a.Example) > 0 {
		return a.Example
	}

	return []interface{}{a.Type.ExampleValue()}
}

type ObjectSchema struct {
	Properties map[string]Schema
	Example    map[string]interface{}
}

func (o ObjectSchema) ExampleValue() interface{} {
	if len(o.Example) > 0 {
		return o.Example
	}

	example := make(map[string]interface{}, len(o.Properties))

	for key, propSchema := range o.Properties {
		example[key] = propSchema.ExampleValue()
	}

	return example
}

type FakerSchema struct {
	Example interface{}
}

func (f FakerSchema) ExampleValue() interface{} {
	return f.Example
}
