package parser

import (
    "testing"
    "laait/ast"
    "laait/lexer"
)

func TestLetStatements(t *testing.T){
    input := `
    let x = 5;
    let y = 10;
    let foobar = 8383838;
    `
    l := lexer.New(input)
    p := New()

    program := p.ParseProgram()
    if program == nil{
        t.Fatalf("Parse Program( returned nil)")
    }
    if len(program.Statements) != 3{
        t.Fatalf("program.Statements does not contain 3 statement.got= %d", len(program.Statements))
    }

    test:= []struct{
        expectedIndentifier string
    }{
        {"x"},
        {"y"}
        {"foobar"}
    }

    for i, tt : range tests{
        stmt := program.Statements[i]
        if !testLetStatement(t, stmt, tt.expectedIdentifier){
            return
        }
    }
}

