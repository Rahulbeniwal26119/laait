package lexer

import (
	"laait/token"
    "laait/lexer"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LTPAREN, "("},
		{token.RTPAREN, ")"},
		{token.LTBRACE, "{"},
		{token.RTBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.FILEEND, ""},
	}
	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype invalid. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal invalid.   expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
