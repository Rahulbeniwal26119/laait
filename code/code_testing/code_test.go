package code

import (
	"laait/code"
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		op       code.Opcode
		operands []int
		expected []byte
	}{
		{code.OpConstant, []int{65534}, []byte{byte(code.OpConstant), 255, 254}},
		{code.OPADD, []int{}, []byte{byte(code.OPADD)}},
		{code.OPGETLOCAL, []int{255}, []byte{byte(code.OPGETLOCAL), 255}},
	}

	for _, tt := range tests {
		instruction := code.Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d",
				len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d",
					i, b, instruction[i])
			}
		}
	}
}

func TestInstructionString(t *testing.T) {
	// instructions := []code.Instructions{
	// 	code.Make(code.OpConstant, 1),
	// 	code.Make(code.OpConstant, 2),
	// 	code.Make(code.OpConstant, 65535),
	// 	code.Make(code.OPADD),
	// }

	instructions := []code.Instructions{
		code.Make(code.OPADD),
		code.Make(code.OPGETLOCAL, 1),
		code.Make(code.OpConstant, 2),
		code.Make(code.OpConstant, 65535),
	}

	// 	expected := `0000 OpConstant 1
	// 0003 OpConstant 2
	// 0006 OpConstant 65535
	// 0009 OPADD
	// `

	expected := `0000 OPADD
0001 OPGETLOCAL 1
0003 OpConstant 2
0006 OpConstant 65535
`

	concatted := code.Instructions{}
	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}

	if concatted.String() != expected {
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot=%q",
			expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        code.Opcode
		operands  []int
		bytesRead int
	}{
		{code.OpConstant, []int{65535}, 2},
		{code.OPGETLOCAL, []int{255}, 1},
	}

	for _, tt := range tests {
		instruction := code.Make(tt.op, tt.operands...)
		def, err := code.Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("definition not found: %q\n", err)
		}

		operandsRead, n := code.ReadOperands(def, instruction[1:])
		if n != tt.bytesRead {
			t.Fatalf("n wrong. want=%d, got=%d", tt.bytesRead, n)
		}
		for i, want := range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("operand wrong. want=%d, got=%d", want, operandsRead[i])
			}
		}
	}
}
