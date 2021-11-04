package ast

import (
	"laait/ast"
	"laait/token"
	"testing"
)

// test the string representation
func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.BIND, Literal: "bind"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.NOUN, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.NOUN, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "bind myVar = anotherVar;" {
		t.Errorf("Program.String() wrong. got %q", program.String())
	}
}
