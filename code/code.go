package code

import (
	"bytes"
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
		case 1:
			Instructions[offset] = byte(o)
		}
		offset += width
	}
	return Instructions
}

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
	OPADD
	OPPOP
	OPSUB
	OPMUL
	OPDIV
	OPTRUE
	OPFALSE
	OPEQUAL
	OPNOTEQUAL
	OPGREATERTHAN
	OPMINUS
	OPBANG
	OPJUMPNOTTRUE
	OPJUMP
	OPNULL
	OPGETGLOBAL
	OPSETGLOBAL
	OPARRAY
	OPHASH
	OPINDEX
	OPCALL
	OPRETURNVALUE
	OPRETURN
	OPGETLOCAL
	OPSETLOCAL
	OPGETBUILTIN
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant:    {"OpConstant", []int{2}},
	OPADD:         {"OPADD", []int{}},
	OPPOP:         {"OPPOP", []int{}},
	OPSUB:         {"OPSUB", []int{}},
	OPMUL:         {"OPMUL", []int{}},
	OPDIV:         {"OPDIV", []int{}},
	OPTRUE:        {"OPTRUE", []int{}},
	OPFALSE:       {"OPFALSE", []int{}},
	OPEQUAL:       {"OPEQUAL", []int{}},
	OPNOTEQUAL:    {"OPNOTEQUAL", []int{}},
	OPGREATERTHAN: {"OPGREATERTHAN", []int{}},
	OPMINUS:       {"OPMINUS", []int{}},
	OPBANG:        {"OPBANG", []int{}},
	OPJUMPNOTTRUE: {"OPJUMPNOTTRUE", []int{2}},
	OPJUMP:        {"OPJUMP", []int{2}},
	OPNULL:        {"OPNULL", []int{}},
	OPGETGLOBAL:   {"OPGETGLOBAL", []int{2}},
	OPSETGLOBAL:   {"OPSETGLOBAL", []int{2}},
	OPARRAY:       {"OPARRAY", []int{2}},
	OPHASH:        {"OPHASH", []int{2}},
	OPINDEX:       {"OPINDEX", []int{}},
	OPCALL:        {"OPCALL", []int{1}}, // store the the count of arguments max is 256
	OPRETURNVALUE: {"OPRETURNVALUE", []int{}},
	OPGETLOCAL:    {"OPGETLOCAL", []int{1}},
	OPSETLOCAL:    {"OPSETLOCAL", []int{1}},
	OPGETBUILTIN:  {"OPGETBUILTIN", []int{1}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		case 1:
			operands[i] = int(ReadUint8(ins[offset:]))

		}
		offset += width
	}
	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

func ReadUint8(ins Instructions) uint8 {
	return uint8(ins[0])
}

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstructions(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstructions(def *Definition, operands []int) string {
	operandsCount := len(def.OperandWidths)

	if len(operands) != operandsCount {
		return fmt.Sprintf("ERROR: operands len %d does not match defined %d\n", len(operands), operandsCount)
	}

	switch operandsCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}
