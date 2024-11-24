package lexer

import "lang/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken is a method on the Lexer struct that returns the next token in the input
// and advances the position of the input to the next token.
// the purpose of the readChar function is to give us the next character and advance our position in the input string.

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal, tok.Type = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '!' || ch == '?'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() (string, token.TokenType) {
	position := l.position
	
	if l.ch == '0' {
		return l.readPrefixedNumber(position)
	}
	
	l.readIntegerPart()
	
	if l.ch == '.' {
		return l.readFloat(position)
	}
	
	return l.input[position:l.position], token.INT
}

func (l *Lexer) readPrefixedNumber(position int) (string, token.TokenType) {
	l.readChar()
	if l.ch == 'x' || l.ch == 'X' {
		return l.readHexNumber()
	} else if l.ch == 'o' || l.ch == 'O' {
		return l.readOctalNumber()
	}
	l.resetPosition(position)
	return "", token.ILLEGAL
}

func (l *Lexer) readHexNumber() (string, token.TokenType) {
	l.readChar()
	start := l.position
	hasDigits := false
	
	for isHexDigit(l.ch) {
		hasDigits = true
		l.readChar()
	}
	
	if !hasDigits {
		return "", token.ILLEGAL
	}
	
	literal := l.input[start:l.position]
	// Check all chars are valid hex
	for _, ch := range literal {
		if !isHexDigit(byte(ch)) {
			return "", token.ILLEGAL
		}
	}
	
	return literal, token.HEX
}

func (l *Lexer) readOctalNumber() (string, token.TokenType) {
	l.readChar()
	start := l.position
	hasDigits := false
	
	for isOctalDigit(l.ch) {
		hasDigits = true
		l.readChar()
	}
	
	if !hasDigits {
		return "", token.ILLEGAL
	}
	
	literal := l.input[start:l.position]
	// Check all chars are valid octal
	for _, ch := range literal {
		if !isOctalDigit(byte(ch)) {
			return "", token.ILLEGAL
		}
	}
	
	return literal, token.OCTAL
}

func (l *Lexer) resetPosition(position int) {
	l.position = position
	l.readPosition = position + 1
	l.ch = l.input[position]
}

func (l *Lexer) readIntegerPart() {
	for isDigit(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readFloat(position int) (string, token.TokenType) {
	l.readChar()
	hasDecimals := false
	for isDigit(l.ch) {
		hasDecimals = true
		l.readChar()
	}
	if hasDecimals {
		return l.input[position:l.position], token.FLOAT
	}
	return l.input[position:l.position], token.INT
}

func isHexDigit(ch byte) bool {
    return isDigit(ch) || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

func isOctalDigit(ch byte) bool {
    return '0' <= ch && ch <= '7'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}