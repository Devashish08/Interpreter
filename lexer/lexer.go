// Package lexer implements a lexical analyzer for processing source code into tokens.

package lexer

import "lang/token"

// Lexer represents the lexical analyzer with the following fields:
//   - input: the source code string to be tokenized
//   - position: current position in input (points to current char)
//   - readPosition: current reading position in input (after current char)
//   - ch: current char under examination
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New creates and initializes a new Lexer with the given input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Initialize by reading the first character
	return l
}

// readChar advances the lexer's position in the input and updates the current character.
// Sets ch to 0 (NULL) when reaching end of input.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken identifies and returns the next token from the input.
// This is the main tokenization function that handles:
// - Operators (+, -, *, /, <, >, =)
// - Delimiters (,;(){}[])
// - Identifiers (variables, keywords)
// - Numbers (integers)
// - Two-character tokens (==, !=)
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
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
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// Helper functions:

// newToken creates a Token with the given type and character
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier reads an identifier from input (letters/underscore)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter checks if character is a letter, underscore, ! or ?
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '!' || ch == '?'
}

// skipWhitespace advances past any whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readNumber reads a numeric sequence from input
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isDigit checks if character is numeric (0-9)
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// peekChar looks at the next character without advancing position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Note: There are commented-out functions for handling:
// - Hexadecimal numbers (0x...)
// - Octal numbers (0o...)
// - Floating point numbers
// These can be uncommented and implemented for extended numeric support
