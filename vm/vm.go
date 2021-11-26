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
		case code.OPADD:
			right := vm.pop()
			left := vm.pop()
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			result := leftValue + rightValue
			vm.push(&object.Integer{Value: result})
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
