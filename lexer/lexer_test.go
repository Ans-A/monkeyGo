package lexer

import (
	"testing"

	"../token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOG, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expcted =%q, got =%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Type != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expcted =%q, got =%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
