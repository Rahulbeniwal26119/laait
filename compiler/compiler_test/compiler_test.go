package compiler

import (
	"fmt"
	"laait/ast"
	"laait/code"
	"laait/compiler"
	"laait/lexer"
	"laait/object"
	"laait/parser"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
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
	concatted := concatInstructionsss(expected)

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
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPADD),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "1; 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPPOP),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "1 - 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPSUB),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "2/1",
			expectedConstants: []interface{}{2, 1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPDIV),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "2 * 3",
			expectedConstants: []interface{}{2, 3},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPMUL),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "-1",
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPMINUS),
				code.Make(code.OPPOP),
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

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed : %s", err)
		}

		err = testConstants(t, tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}

func concatInstructionsss(s []code.Instructions) code.Instructions {
	out := code.Instructions{}

	for _, ins := range s {
		out = append(out, ins...)
	}

	return out
}

func testConstants(
	t *testing.T,
	expected []interface{},
	actual []object.Object,
) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants. got=%d, want=%d",
			len(actual), len(expected))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			err := testIntegerObject(int64(constant), actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testIntegerObject failed : %s", i, err)
			}

		case string:
			err := testStringObject(constant, actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testStringObject failed: %s", i, err)
			}

		case []code.Instructions:
			fn, ok := actual[i].(*object.CompiledFunction)
			if !ok {
				return fmt.Errorf("constant %d - not a function : %T",
					i, actual[i])
			}

			err := testInstructions(constant, fn.Instructions)
			if err != nil {
				return fmt.Errorf("constant %d - testInstructions failed: %s",
					i, err)
			}
		}
	}

	return nil
}

func testStringObject(expected string, actual object.Object) error {
	result, ok := actual.(*object.String)
	if !ok {
		return fmt.Errorf("object is not String. got %T (+%v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%q, want=%q", result.Value, expected)
	}

	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)",
			actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
	}

	return nil
}

func TestBooleanExpressions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "true",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OPTRUE),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "false",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OPFALSE),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "1 > 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPGREATERTHAN),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "1 < 2",
			expectedConstants: []interface{}{2, 1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPGREATERTHAN),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "1 == 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPEQUAL),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "1 != 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPNOTEQUAL),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "true == false",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OPTRUE),
				code.Make(code.OPFALSE),
				code.Make(code.OPEQUAL),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "true != false",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OPTRUE),
				code.Make(code.OPFALSE),
				code.Make(code.OPNOTEQUAL),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "!true",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OPTRUE),
				code.Make(code.OPBANG),
				code.Make(code.OPPOP),
			},
		},
	}
	runCompilerTest(t, tests)
}

func TestConditionals(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
			if (true) { 10 } else { 20 }; 3333;
			`,
			expectedConstants: []interface{}{10, 20, 3333},
			expectedInstructions: []code.Instructions{
				// 0000
				code.Make(code.OPTRUE),
				// 0001
				code.Make(code.OPJUMPNOTTRUE, 10),
				// 0004
				code.Make(code.OpConstant, 0),
				// 0007
				code.Make(code.OPJUMP, 13),
				// 0010
				code.Make(code.OpConstant, 1),
				// 0013
				code.Make(code.OPPOP),
				// 0014
				code.Make(code.OpConstant, 2),
				// 0017
				code.Make(code.OPPOP),
			},
		},
		{
			input: `
			if (true) { 10}; 3333;
			`,
			expectedConstants: []interface{}{10, 3333},
			expectedInstructions: []code.Instructions{
				// 0000
				code.Make(code.OPTRUE),
				// 0001
				code.Make(code.OPJUMPNOTTRUE, 10),
				// 0004
				code.Make(code.OpConstant, 0),
				// 0007
				code.Make(code.OPJUMP, 11),
				// 0010
				code.Make(code.OPNULL),
				// 0011
				code.Make(code.OPPOP),
				// 0012
				code.Make(code.OpConstant, 1),
				// 0015
				code.Make(code.OPPOP),
			},
		},
	}

	runCompilerTest(t, tests)
}

func TestGlobalLetStatements(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
			let one = 1;
			let two = 2;
			`,
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPSETGLOBAL, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPSETGLOBAL, 1),
			},
		},
		{
			input: `
			let one = 1;
			one;
			`,
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPSETGLOBAL, 0),
				code.Make(code.OPGETGLOBAL, 0),
				code.Make(code.OPPOP),
			},
		},
		{
			input: `
			let one = 1;
			let two = one;
			two;
			`,
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPSETGLOBAL, 0),
				code.Make(code.OPGETGLOBAL, 0),
				code.Make(code.OPSETGLOBAL, 1),
				code.Make(code.OPGETGLOBAL, 1),
				code.Make(code.OPPOP),
			},
		},
	}
	runCompilerTest(t, tests)
}

func TestStringExpressions(t *testing.T) {
	tests := []compilerTestCase{
		{input: `"laait"`,
			expectedConstants: []interface{}{"laait"},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPPOP)},
		},
		{
			input:             `"la" + "ait"`,
			expectedConstants: []interface{}{"la", "ait"},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPADD),
				code.Make(code.OPPOP),
			},
		},
	}

	runCompilerTest(t, tests)
}

func TestArrayLiterals(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "[]",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OPARRAY, 0),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "[1,2,3]",
			expectedConstants: []interface{}{1, 2, 3},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OPARRAY, 3),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "[1+2, 3 -4 , 5 * 6]",
			expectedConstants: []interface{}{1, 2, 3, 4, 5, 6},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPADD),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OPSUB),
				code.Make(code.OpConstant, 4),
				code.Make(code.OpConstant, 5),
				code.Make(code.OPMUL),
				code.Make(code.OPARRAY, 3),
				code.Make(code.OPPOP),
			},
		},
	}
	runCompilerTest(t, tests)
}

func TestHashLiterals(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "{}",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OPHASH, 0),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "{1 : 2, 3 : 4, 5 :6}",
			expectedConstants: []interface{}{1, 2, 3, 4, 5, 6},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpConstant, 4),
				code.Make(code.OpConstant, 5),
				code.Make(code.OPHASH, 6),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "{1: 2+3, 4 : 5*6}",
			expectedConstants: []interface{}{1, 2, 3, 4, 5, 6},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OPADD),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpConstant, 4),
				code.Make(code.OpConstant, 5),
				code.Make(code.OPMUL),
				code.Make(code.OPHASH, 4),
				code.Make(code.OPPOP),
			},
		},
	}

	runCompilerTest(t, tests)
}

func TestIndexExpression(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "[1,2,3][1+1]",
			expectedConstants: []interface{}{1, 2, 3, 1, 1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OPARRAY, 3),
				code.Make(code.OpConstant, 3),
				code.Make(code.OpConstant, 4),
				code.Make(code.OPADD),
				code.Make(code.OPINDEX),
				code.Make(code.OPPOP),
			},
		},
		{
			input:             "{1 : 2}[2-1]",
			expectedConstants: []interface{}{1, 2, 2, 1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPHASH, 2),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OPSUB),
				code.Make(code.OPINDEX),
				code.Make(code.OPPOP),
			},
		},
	}
	runCompilerTest(t, tests)
}

func TestFunctions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `function() { return 5 + 10 }`,
			expectedConstants: []interface{}{
				5,
				10,
				[]code.Instructions{
					code.Make(code.OpConstant, 0),
					code.Make(code.OpConstant, 1),
					code.Make(code.OPADD),
					code.Make(code.OPRETURNVALUE),
				},
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 2),
				code.Make(code.OPPOP),
			},
		},
		{
			input: `function() { 24 }();`,
			expectedConstants: []interface{}{
				24,
				[]code.Instructions{
					code.Make(code.OpConstant, 0),
					code.Make(code.OPRETURNVALUE),
				},
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 1),
				code.Make(code.OPCALL),
				code.Make(code.OPPOP),
			},
		},
		{
			input: `let noArgs = function() { 24 };
					noArgs();`,
			expectedConstants: []interface{}{
				24,
				[]code.Instructions{
					code.Make(code.OpConstant, 0),
					code.Make(code.OPRETURNVALUE),
				},
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 1),
				code.Make(code.OPSETGLOBAL, 0),
				code.Make(code.OPGETGLOBAL, 0),
				code.Make(code.OPCALL),
				code.Make(code.OPPOP),
			},
		},
	}
	runCompilerTest(t, tests)
}

func TestCompilerWithoutReturnValue(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `function() {}`,
			expectedConstants: []interface{}{
				[]code.Instructions{
					code.Make(code.OPRETURN),
				},
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPPOP),
			},
		},
	}

	runCompilerTest(t, tests)
}

func TestLetStatementScopes(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
			let num  = 55;
			function() { num };
			`,
			expectedConstants: []interface{}{
				55,
				[]code.Instructions{
					code.Make(code.OPGETGLOBAL, 0),
					code.Make(code.OPRETURNVALUE),
				},
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPSETGLOBAL, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPPOP),
			},
		},
		{
			input: `
			function() {
				let num = 55;
				num
			}`,
			expectedConstants: []interface{}{
				55,
				[]code.Instructions{
					code.Make(code.OpConstant, 0),
					code.Make(code.OPSETLOCAL, 0),
					code.Make(code.OPGETLOCAL, 0),
					code.Make(code.OPRETURNVALUE),
				},
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 1),
				code.Make(code.OPPOP),
			},
		},
		{
			input: `
					function() {
						let a = 55;
						let b = 77;
						a + b
					}`,
			expectedConstants: []interface{}{
				55,
				77,
				[]code.Instructions{
					code.Make(code.OpConstant, 0),
					code.Make(code.OPSETLOCAL, 0),
					code.Make(code.OpConstant, 1),
					code.Make(code.OPSETLOCAL, 1),
					code.Make(code.OPGETLOCAL, 0),
					code.Make(code.OPGETLOCAL, 1),
					code.Make(code.OPADD),
					code.Make(code.OPRETURNVALUE),
				},
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 2),
				code.Make(code.OPPOP),
			},
		},
	}
	runCompilerTest(t, tests)
}

// func TestCompilerScopes(t *testing.T) {
// 	c := compiler.New()
// 	if c.scopeIndex != 0 {
// 		t.Errorf("scopeIndex wrong. got=%d, want=%d", compiler.scopeIndex, 0)
// 	}

// 	globalSymbolTable := c.SymbolTable
// 	c.emit(code.OPMUL)

// 	c.enterScope()

// 	if c.scopeIndex != 1 {
// 		t.Errorf("scopeIndex wrong. got=%d, want=%d", c.scopeIndex, 1)
// 	}

// 	c.emit(scope.OPSUB)
// 	if len(c.scopes[c.scopeIndex].instructions) != 1 {
// 		t.Errof("instructions length wrong. got=%d", len(c.scopes[c.scopeIndex].instructions))
// 	}

// 	last := c.scopes[c.scopeIndex].LastInstructions
// 	if last.Opcode != code.OPSUB {
// 		t.Errof("lastInstruction.Opcode wrong. got=%d, want=%d", last.Opcode, code.OPSUB)
// 	}

// 	if c.symbolTable.Outer != globalSymbolTable {
// 		t.Errorf("compiler didnot not enclode symbol table")
// 	}

// 	c.leaveScope()
// 	if c.scopeIndex != 0 {
// 		t.Errof("scopeIndex wrong. got=%d, want=%d", c.scopeIndex, 0)
// 	}

// 	if c.symbolTable != globalSymbolTable {
// 		t.Errorf("compiler didnot restore global symbol table")
// 	}

// 	if c.symbolTable.Outer != nil {
// 		t.Errorf("compiler modified global symbol table incorrectly.")
// 	}

// 	c.emit(code.OPADD)

// 	if len(compiler.scopes[compiler.scopeIndex].instructions) != 2 {
// 		t.Errorf("instructions length wrong. got=%d",
// 			len(c.scopes[c.scopeIndex].instructions))
// 	}
// 	last = compiler.scopes[compiler.scopeIndex].lastInstruction
// 	if last.Opcode != code.OPADD {
// 		t.Errof("lastInstruction.Opcode wrong. got=%d, want=%d",
// 			last.Opcode, code.OPADD)
// 	}

// 	previous := c.scopes[c.scopeIndex].previousInstruction
// 	if previous.Opcode != code.OPMUL {
// 		t.Errorf("previousInstruction.Opcode wrong. got=%d, want=%d",
// 			previous.Opcode, code.OPMUL)
// 	}

// }

func TestFunctionsCalls(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
			let oneArg = function(a) { };
			oneArg(24);
			`,
			expectedConstants: []interface{}{
				[]code.Instructions{
					code.Make(code.OPRETURN),
				},
				24,
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPSETGLOBAL, 0),
				code.Make(code.OPGETGLOBAL, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OPCALL, 1),
				code.Make(code.OPPOP),
			},
		},
		{
			input: `
			let manyArgs = function(a, b, c){ };
			manyArgs(24, 25, 26);
			`,
			expectedConstants: []interface{}{
				[]code.Instructions{
					code.Make(code.OPRETURN),
				},
				24,
				25,
				26,
			},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OPSETGLOBAL, 0),
				code.Make(code.OPGETGLOBAL, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpConstant, 2),
				code.Make(code.OpConstant, 3),
				code.Make(code.OPCALL, 3),
				code.Make(code.OPPOP),
			},
		},
	}
	runCompilerTest(t, tests)
}
