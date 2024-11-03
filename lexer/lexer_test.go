package lexer

import (
	"mrvi_interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `MrviFn{}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Function, "MrviFn"},
		{token.LBrace, "{"},
		{token.RBrace, "}"},
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
