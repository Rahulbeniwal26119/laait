package compiler

import (
	"fmt"
	"laait/ast"
	"laait/code"
	"laait/compiler"
	"laait/lexer"
	"laait/parser"
	"testing"
)

type compilerTestCase struct {
	input               string
	expectedConstants   []interface{}
	expectedInstruction []code.Instructions
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testInstructions(
	expected []code.Instructions,
	actual code.Instructions,
) error {
	concatted := concatInstructions(expected)

	if len(actual) != len(concatted) {
		return fmt.Errorf("wrong instruction length.\nwant=%q\ngot =%q",
			concatted, actual)
	}

	for i, ins := range concatted {
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot=%q",
				i, concatted, actual)
		}
	}
	return nil
}
func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstruction: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			},
		},
	}

	runCompilerTest(t, tests)
}

func runCompilerTest(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		compiler := compiler.New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("compiler error : %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstruction(tt.expectedInstruction, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstruction failed : %s", err)
		}

		err = testConstants(tt, tt.expectedConstants, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}
