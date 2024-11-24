package lexer

import (
	"testing"

	"lang/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	  x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if(5 < 10) {
return true;
} else {
 return false;
 }

    10 == 10;
    10 != 9;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// func TestAdvancedNumberFormats(t *testing.T) {
// 	input := `
//         let hex = 0xFF;
//         let oct = 0o77;
//         let float = 123.456;
//         let plain = 42;
//     `

// 	tests := []struct {
// 		expectedType    token.TokenType
// 		expectedLiteral string
// 	}{
// 		{token.LET, "let"},
// 		{token.IDENT, "hex"},
// 		{token.ASSIGN, "="},
// 		{token.HEX, "FF"},
// 		{token.SEMICOLON, ";"},

// 		{token.LET, "let"},
// 		{token.IDENT, "oct"},
// 		{token.ASSIGN, "="},
// 		{token.OCTAL, "77"},
// 		{token.SEMICOLON, ";"},

// 		{token.LET, "let"},
// 		{token.IDENT, "float"},
// 		{token.ASSIGN, "="},
// 		{token.FLOAT, "123.456"},
// 		{token.SEMICOLON, ";"},

// 		{token.LET, "let"},
// 		{token.IDENT, "plain"},
// 		{token.ASSIGN, "="},
// 		{token.INT, "42"},
// 		{token.SEMICOLON, ";"},

// 		{token.EOF, ""},
// 	}

// 	l := New(input)

// 	for i, tt := range tests {
// 		tok := l.NextToken()

// 		if tok.Type != tt.expectedType {
// 			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
// 				i, tt.expectedType, tok.Type)
// 		}

// 		if tok.Literal != tt.expectedLiteral {
// 			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
// 				i, tt.expectedLiteral, tok.Literal)
// 		}
// 	}
// }

// func TestInvalidNumbers(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected token.TokenType
// 	}{
// 		{"0x", token.ILLEGAL},  // Incomplete hex
// 		{"0xG", token.ILLEGAL}, // Invalid hex digit
// 		{"0o8", token.ILLEGAL}, // Invalid octal digit
// 		{"0o", token.ILLEGAL},  // Incomplete octal
// 		{"123.", token.INT},    // Dot without decimals
// 	}

// 	for i, tt := range tests {
// 		l := New(tt.input)
// 		tok := l.NextToken()

// 		if tok.Type != tt.expected {
// 			t.Errorf("tests[%d] - tokentype wrong. input=%q, expected=%q, got=%q",
// 				i, tt.input, tt.expected, tok.Type)
// 		}
// 	}
// }
