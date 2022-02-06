package parser

import (
	"fmt"
	"testing"
)

func TestLexer_SimpleType(t *testing.T) {
	schemaDefinition := `
{
  hero {
    name
  }
}
`

	l := NewLexer(schemaDefinition)

	tests := []struct {
		tokenType    TokenType
		tokenLiteral string
	}{
		{LBRACE, "{"},
		{IDENT, "hero"},
		{LBRACE, "{"},
		{IDENT, "name"},
		{RBRACE, "}"},
		{RBRACE, "}"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s (%s)", tt.tokenLiteral, TokenTypeName(tt.tokenType))

		t.Run(name, func(t *testing.T) {
			token := l.NextToken()

			if token.Type != tt.tokenType {
				t.Fatalf("unexpected TokenType - want: %s, got: %s", TokenTypeName(tt.tokenType), TokenTypeName(token.Type))
			}
			if token.Literal != tt.tokenLiteral {
				t.Fatalf("unexpected Literal - want: %s, got: %s", tt.tokenLiteral, token.Literal)
			}
		})
	}
}
