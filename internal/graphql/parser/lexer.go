package parser

type Lexer struct {
	input        string
	position     int
	ch           byte
	readPosition int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.read()
	return l
}

func (l *Lexer) NextToken() Token {
	if l.isEOF() {
		return Token{Type: EOF}
	}

	l.skipWhitespace()

	var token Token
	switch l.ch {
	case '{':
		token = Token{Type: LBRACE, Literal: "{"}
	case '}':
		token = Token{Type: RBRACE, Literal: "}"}
	case ':':
		token = Token{Type: COLON, Literal: ":"}
	case 0:
		return Token{Type: EOF}
	default:
		if !isLetter(l.ch) {
			token = Token{Type: ILLEGAL, Literal: string(l.ch)}
			break
		}

		ident := l.readIdent()
		if ident == "type" {
			token = Token{Type: TYPE, Literal: ident}
		} else {
			token = Token{Type: IDENT, Literal: ident}
		}

	}

	l.read()

	return token
}

func (l *Lexer) isEOF() bool {
	return l.readPosition >= len(l.input)
}

func (l *Lexer) read() {
	if l.isEOF() {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peek() byte {
	return l.input[l.readPosition]
}

func (l *Lexer) readIdent() string {
	start := l.position
	for isLetter(l.peek()) {
		l.read()
	}

	return l.input[start:l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.read()
	}
}

func isLetter(ch byte) bool {
	if 'a' <= ch && ch <= 'z' {
		return true
	}
	if 'A' <= ch && ch <= 'Z' {
		return true
	}
	return false
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
