package lexer_test

import (
	"laait/lexer"
	"laait/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input1 := `=+(){},;`
	input2 := `
    bind five = 5;
    bind ten = 10;
    bind add = fn(x,y){
        x + y;
    };
    bind result = add(five, ten);
    !-/*5;
    5 < 10 > 5;
	10 == 10
	10 != 9
    `

	// var Token struct {
	// 	expectedType    token.TokenType
	// 	expectedLiteral string
	// }

	tests1 := []struct {
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

	tests2 := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BIND, "bind"},
		{token.NOUN, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.BIND, "bind"},
		{token.NOUN, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.BIND, "bind"},
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
		{token.BIND, "bind"},
		{token.NOUN, "result"},
		{token.ASSIGN, "="},
		{token.NOUN, "add"},
		{token.LTPAREN, "("},
		{token.NOUN, "five"},
		{token.COMMA, ","},
		{token.NOUN, "ten"},
		{token.RTPAREN, ")"},
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
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.FILEEND, ""},
	}

	l1 := lexer.New(input1)
	l2 := lexer.New(input2)

	testCode(t, l1, tests1)
	testCode(t, l2, tests2)

}

func testCode(t *testing.T, l *lexer.Lexer, test []struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
	for i, tt := range test {
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
