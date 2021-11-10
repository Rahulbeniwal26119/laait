package code

import (
	"encoding/binary"
	"fmt"
)

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	InstructionsLen := 1
	for _, w := range def.OperandWidths {
		InstructionsLen += w
	}

	Instructions := make([]byte, InstructionsLen)
	Instructions[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(Instructions[offset:], uint16(o))
		}
		offset += width
	}
	return Instructions
}

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}
