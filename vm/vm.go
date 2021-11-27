// This Virtual Machine is a stack based VM instead of a register based VM
package vm

import (
	"fmt"
	"laait/code"
	"laait/compiler"
	"laait/object"
)

const StackSize = 2560

type VM struct {
	constant     []object.Object
	instructions code.Instructions

	stack []object.Object
	sp    int // Always point the upcoming value
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constant:     bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	// fetch part
	// fetch - decode - execute loop
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2
			err := vm.push(vm.constant[constIndex])
			if err != nil {
				return err
			}
		case code.OPADD, code.OPSUB, code.OPMUL, code.OPDIV:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}
		case code.OPPOP:
			vm.pop()
		}
	}

	return nil
}

func (vm *VM) push(object object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("Stack Overflow")
	}

	vm.stack[vm.sp] = object
	vm.sp++
	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}

func (vm *VM) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	leftType := left.Type()
	rightType := right.Type()

	if leftType == object.INTEGER_OBJ && rightType == object.INTEGER_OBJ {
		return vm.executeBinaryIntergerOperation(op, left, right)
	}

	return fmt.Errorf("unsupported types for binary operation: %s %s", leftType, rightType)
}

func (vm *VM) executeBinaryIntergerOperation(op code.Opcode, left, right object.Object) error {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	var result int64

	switch op {
	case code.OPADD:
		result = leftValue + rightValue
	case code.OPSUB:
		result = leftValue - rightValue
	case code.OPMUL:
		result = leftValue * rightValue
	case code.OPDIV:
		result = leftValue / rightValue
	default:
		return fmt.Errorf("unknown integer operator: %d", op)
	}

	return vm.push(&object.Integer{Value: result})
}
