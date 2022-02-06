package model

type TypeSystem struct {
	Objects []ObjectType
}

func (ts TypeSystem) ObjectTypeByName(name string) (ObjectType, bool) {
	for _, o := range ts.Objects {
		if o.Name == name {
			return o, true
		}
	}

	return ObjectType{}, false
}

type ObjectType struct {
	Name   string
	Fields []Field
}

func (t ObjectType) FieldByName(name string) (Field, bool) {
	for _, f := range t.Fields {
		if f.Name == name {
			return f, true
		}
	}

	return Field{}, false
}

type Field struct {
	Name string
	Type string
}
