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

Func TestNextToken(t *testing.T){
    input := `
    mini five = 5;
    mini ten = 10;
    mini add = fn(x,y){
        x + y;
    };
    mini result = add(five, ten);
    `
    tests := []struct{
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.MINI, "mini"},
        {token.NOUN, "five"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.MINI, "let"},
        {token.NOUN, "ten"},
        {token.ASSIGN, "="}
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.MINI, "mini"},
        {token.NOUN, "add"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "fn"},
        {token.LTPAREN, "("},
        {token.NOUN, "x"},
        {token.COMMA, ","},
        {token.NOUN, "y"},
        {token.RTPAREN, ")"},
        {token.LTBRACE, "{"},
        {token.NOUN, "x"},
        {token.PLUS, "+"},
        {token.NOUN, "y"},
        {token.SEMICOLON, ";"},
        {token.RTBRACE, "}"},
        {token.SEMICOLON, ";"},
        {token.MINI, "mini"},
        {token.NOUN, "result"},
        {token.ASSIGN, "="},
        {token.NOUN, "add"},
        {token.LTPAREN, "("},
        {token.NOUN, "five"},
        {token.COMMA, ","},
        {token.NOUN, "ten"},
        {token.RTPAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.FILEEND, ""}
    }
}
