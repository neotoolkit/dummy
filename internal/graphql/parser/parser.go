package parser

import "github.com/neotoolkit/dummy/internal/graphql/model"

type QueryParser struct {
	l *Lexer

	curToken  Token
	peekToken Token
}

func NewQueryParser(l *Lexer) *QueryParser {
	p := &QueryParser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *QueryParser) Parse() *model.Query {
	q := &model.Query{}

	for !p.isEOF() {
		switch p.curToken.Type {
		case LBRACE:
			q.SelectionSet = p.parseSelectionSet()
		}

		p.nextToken()
	}

	return q
}

func (p *QueryParser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *QueryParser) isEOF() bool {
	return p.curToken.Type == EOF
}

func (p *QueryParser) parseSelectionSet() model.SelectionSet {
	var (
		selectionSet model.SelectionSet
	)

	for p.peekToken.Type != RBRACE {
		p.nextToken()

		switch {
		case p.curToken.Type == IDENT:
			selectionSet = append(selectionSet, model.SelectionField{Name: p.curToken.Literal})
		case p.curToken.Type == LBRACE:
			selectionSet[len(selectionSet)-1].SelectionSet = p.parseSelectionSet()
		default:
			return nil
		}
	}

	p.nextToken()

	return selectionSet
}

type TypeSystemParser struct {
	l *Lexer

	curToken  Token
	peekToken Token
}

func NewTypeSystemParser(l *Lexer) *TypeSystemParser {
	p := &TypeSystemParser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *TypeSystemParser) Parse() *model.TypeSystem {
	ts := &model.TypeSystem{}

	for !p.isEOF() {
		switch p.curToken.Type {
		case TYPE:
			typeDef := p.parseTypeDefinition()
			if typeDef != nil {
				ts.Objects = append(ts.Objects, *typeDef)
			}
		}

		p.nextToken()
	}

	return ts
}

func (p *TypeSystemParser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *TypeSystemParser) isEOF() bool {
	return p.curToken.Type == EOF
}

func (p *TypeSystemParser) parseTypeDefinition() *model.ObjectType {
	if p.peekToken.Type != IDENT {
		return nil
	}

	typeDefName := p.peekToken.Literal

	p.nextToken()
	if p.peekToken.Type != LBRACE {
		return nil
	}

	var fields []model.Field
	for p.peekToken.Type != RBRACE {
		p.nextToken()

		field := p.parseFieldDefinition()
		if field != nil {
			fields = append(fields, *field)
		}
	}

	p.nextToken()

	return &model.ObjectType{
		Name:   typeDefName,
		Fields: fields,
	}
}

func (p *TypeSystemParser) parseFieldDefinition() *model.Field {
	if p.curToken.Type != IDENT {
		return nil
	}

	fieldDefName := p.curToken.Literal

	if p.peekToken.Type != COLON {
		return nil
	}

	p.nextToken()
	if p.peekToken.Type != IDENT {
		return nil
	}

	p.nextToken()
	fieldDefType := p.curToken.Literal

	return &model.Field{
		Name: fieldDefName,
		Type: fieldDefType,
	}
}
