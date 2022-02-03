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
	Code      int
	MediaType string
	Schema    Schema
	Example   interface{}
	Examples  map[string]interface{}
}

func (r Response) StatusCode() int {
	return r.Code
}

// ExampleValue -.
func (r Response) ExampleValue(key string) interface{} {
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
	ExampleValue() interface{}
}

// BooleanSchema -.
type BooleanSchema struct {
	Example bool
}

// ExampleValue -.
func (b BooleanSchema) ExampleValue() interface{} {
	return b.Example
}

// IntSchema -.
type IntSchema struct {
	Example int64
}

// ExampleValue -.
func (i IntSchema) ExampleValue() interface{} {
	return i.Example
}

// FloatSchema -.
type FloatSchema struct {
	Example float64
}

// ExampleValue -.
func (f FloatSchema) ExampleValue() interface{} {
	return f.Example
}

// StringSchema -.
type StringSchema struct {
	Example string
}

// ExampleValue -.
func (s StringSchema) ExampleValue() interface{} {
	return s.Example
}

// ArraySchema -.
type ArraySchema struct {
	Type    Schema
	Example []interface{}
}

// ExampleValue -.
func (a ArraySchema) ExampleValue() interface{} {
	if len(a.Example) > 0 {
		return a.Example
	}

	return []interface{}{a.Type.ExampleValue()}
}

// ObjectSchema -.
type ObjectSchema struct {
	Properties map[string]Schema
	Example    map[string]interface{}
}

// ExampleValue -.
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

// FakerSchema -.
type FakerSchema struct {
	Example interface{}
}

// ExampleValue -.
func (f FakerSchema) ExampleValue() interface{} {
	return f.Example
}
