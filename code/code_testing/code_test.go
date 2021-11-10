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
