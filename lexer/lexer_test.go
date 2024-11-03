package lexer

import (
	"mrvi_interpreter/token"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `MrviFn{{b}}( a = 5; b = 10 ; MrviReturn a+b;)`
	input = strings.ReplaceAll(input, " ", "")
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Function, "MrviFn"},
		{token.LBrace, "{"},
		{token.LBrace, "{"},
		{token.Identifiers, "b"},
		{token.RBrace, "}"},
		{token.RBrace, "}"},
		{token.LParen, "("},
		{token.Identifiers, "a"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Identifiers, "b"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Return, "MrviReturn"},
		{token.Identifiers, "a"},
		{token.Plus, "+"},
		{token.Identifiers, "b"},
		{token.Semicolon, ";"},
		{token.RParen, ")"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
