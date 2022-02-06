package parser

type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF

	IDENT

	TYPE

	LBRACE
	RBRACE
	COLON
)

var typeNames = []string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	IDENT:   "IDENT",
	TYPE:    "TYPE",
	LBRACE:  "{",
	RBRACE:  "}",
	COLON:   ":",
}

func TokenTypeName(tt TokenType) string {
	return typeNames[tt]
}

type Token struct {
	Type    TokenType
	Literal string
}
