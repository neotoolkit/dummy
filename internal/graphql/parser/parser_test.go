package parser

import (
	"reflect"
	"testing"

	"github.com/neotoolkit/dummy/internal/graphql/model"
)

func TestQueryParser_Parse(t *testing.T) {
	simpleSchema := `
{ 
  hero { 
    name 
  } 
}`

	tests := []struct {
		name  string
		input string
		want  *model.Query
	}{
		{
			name:  "simple schema",
			input: simpleSchema,
			want: &model.Query{
				SelectionSet: []model.SelectionField{
					{Name: "hero", SelectionSet: []model.SelectionField{
						{Name: "name"},
					}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			p := NewQueryParser(l)

			if got := p.Parse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeSystemParser_Parse(t *testing.T) {
	simpleSchema := `
type Query {
  hero: Character
}

type Character {
  id: Int
  name: String
}
`

	tests := []struct {
		name  string
		input string
		want  *model.TypeSystem
	}{
		{
			name:  "simple schema",
			input: simpleSchema,
			want: &model.TypeSystem{
				Objects: []model.ObjectType{
					{
						Name: "Query",
						Fields: []model.Field{
							{Name: "hero", Type: "Character"},
						},
					},
					{
						Name: "Character",
						Fields: []model.Field{
							{Name: "id", Type: "Int"},
							{Name: "name", Type: "String"},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			p := NewTypeSystemParser(l)

			if got := p.Parse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
