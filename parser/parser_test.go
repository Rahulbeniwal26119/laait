package parser_test

import (
	"fmt"
	"laait/ast"
	"laait/lexer"
	"laait/parser"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    bind  y = 10;
    let  foobar = 1221221;
    `
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Parse Program( returned nil)")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statement.got= %d", len(program.Statements))
	}

	tests := []struct {
		expectedIndentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIndentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	fmt.Println(s.TokenLiteral() == "let")
	if s.TokenLiteral() != "let" && s.TokenLiteral() != "bind" {
		t.Errorf("s.TokenLiteral not 'let' or 'bind'. got = %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}
