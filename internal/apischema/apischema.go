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
	Example    any
	Examples   map[string]any
}

func (r Response) ExampleValue(key string) any {
	if r.Schema == nil {
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
	ExampleValue() any
}

type BooleanSchema struct {
	Example bool
}

func (b BooleanSchema) ExampleValue() any {
	return b.Example
}

type IntSchema struct {
	Example int64
}

func (i IntSchema) ExampleValue() any {
	return i.Example
}

type FloatSchema struct {
	Example float64
}

func (f FloatSchema) ExampleValue() any {
	return f.Example
}

type StringSchema struct {
	Example string
}

func (s StringSchema) ExampleValue() any {
	return s.Example
}

type ArraySchema struct {
	Type    Schema
	Example []any
}

func (a ArraySchema) ExampleValue() any {
	if len(a.Example) > 0 {
		return a.Example
	}

	return []any{a.Type.ExampleValue()}
}

type ObjectSchema struct {
	Properties map[string]Schema
	Example    map[string]any
}

func (o ObjectSchema) ExampleValue() any {
	if len(o.Example) > 0 {
		return o.Example
	}

	example := make(map[string]any, len(o.Properties))

	for key, propSchema := range o.Properties {
		example[key] = propSchema.ExampleValue()
	}

	return example
}

type FakerSchema struct {
	Example any
}

func (f FakerSchema) ExampleValue() any {
	return f.Example
}
