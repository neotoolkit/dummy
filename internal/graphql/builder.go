package graphql

import (
	gModel "github.com/neotoolkit/dummy/internal/graphql/model"
	"github.com/neotoolkit/dummy/internal/graphql/parser"
	"github.com/neotoolkit/dummy/internal/model"
)

func NewBuilder(schemaData []byte) *Builder {
	return &Builder{schemaData: schemaData}
}

type Builder struct {
	schemaData []byte
}

func (b *Builder) Build() (model.API, error) {
	p := parser.NewTypeSystemParser(parser.NewLexer(string(b.schemaData)))
	typeSystem := p.Parse()

	var rootObject gModel.ObjectType
	for _, obj := range typeSystem.Objects {
		if obj.Name == "Query" {
			rootObject = obj
			break
		}
	}

	return API{root: rootObject, ts: typeSystem}, nil
}
