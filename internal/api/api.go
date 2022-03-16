package api

// API -.
type API struct {
	Operations []Operation
}

// Operation -.
type Operation struct {
	Method    string
	Path      string
	Body      map[string]FieldType
	Responses []Response
}

// FieldType -.
type FieldType struct {
	Required bool
	Type     string
}

// Response -.
type Response struct {
	StatusCode int
	MediaType  string
	Schema     Schema
	Example    any
	Examples   map[string]any
}

// ExampleValue -.
func (r Response) ExampleValue(key string) any {
	if nil == r.Schema {
		return nil
	}

	example, ok := r.Examples[key]
	if ok {
		return example
	}

	if r.Example != nil {
		return r.Example
	}

	return r.Schema.ExampleValue()
}

// Schema -.
type Schema interface {
	ExampleValue() any
}

// BooleanSchema -.
type BooleanSchema struct {
	Example bool
}

// ExampleValue -.
func (b BooleanSchema) ExampleValue() any {
	return b.Example
}

// IntSchema -.
type IntSchema struct {
	Example int64
}

// ExampleValue -.
func (i IntSchema) ExampleValue() any {
	return i.Example
}

// FloatSchema -.
type FloatSchema struct {
	Example float64
}

// ExampleValue -.
func (f FloatSchema) ExampleValue() any {
	return f.Example
}

// StringSchema -.
type StringSchema struct {
	Example string
}

// ExampleValue -.
func (s StringSchema) ExampleValue() any {
	return s.Example
}

// ArraySchema -.
type ArraySchema struct {
	Type    Schema
	Example []interface{}
}

// ExampleValue -.
func (a ArraySchema) ExampleValue() any {
	if len(a.Example) > 0 {
		return a.Example
	}

	return []any{a.Type.ExampleValue()}
}

// ObjectSchema -.
type ObjectSchema struct {
	Properties map[string]Schema
	Example    map[string]any
}

// ExampleValue -.
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

// FakerSchema -.
type FakerSchema struct {
	Example any
}

// ExampleValue -.
func (f FakerSchema) ExampleValue() any {
	return f.Example
}
