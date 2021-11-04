package evaluator

import (
	"laait/ast"
	"laait/environment"
	"laait/evaluator"
	"laait/lexer"
	"laait/parser"
	"testing"
)

func TestDefineMacros(t *testing.T) {
	input := `
	let number = 1;
	let fn = function(x, y) { x + y };
	let mymacro = macro(x, y) { x + y };`

	env := environment.NewEnvironment()
	program := testParseProgram(input)

	evaluator.DefineMacros(program, env)

	if len(program.Statements) != 2 {
		t.Fatalf("Wrong number of statements. got=%d", len(program.Statements))
	}

	_, ok := env.Get("number")
	if ok {
		t.Fatalf("number should not be defined")
	}

	_, ok = env.Get("function")
	if ok {
		t.Fatalf("function should not be defined")
	}

	obj, ok := env.Get("mymacro")
	if !ok {
		t.Fatalf("macro not in environment")
	}

	macro, ok := obj.(*environment.Macro)
	if !ok {
		t.Fatalf("Wrong number of macro parameters. got=%d", len(macro.Parameters))
	}

	if macro.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", macro.Parameters[0])
	}

	if macro.Parameters[1].String() != "y" {
		t.Fatalf("parameter is not 'y'. got=%q", macro.Parameters[1])
	}

	expectedBody := "(x + y)"
	if macro.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, macro.Body.String())
	}
}

func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
