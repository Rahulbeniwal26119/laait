// This Virtual Machine is a stack based VM instead of a register based VM
package vm

import (
	"fmt"
	"laait/code"
	"laait/compiler"
	"laait/object"
)

const StackSize = 2560
const GlobalSize = 65536

var True = &object.Boolean{Value: true}
var False = &object.Boolean{Value: false}
var Null = &object.Null{}

type VM struct {
	constant     []object.Object
	instructions code.Instructions

	stack   []object.Object
	sp      int // Always point the upcoming value
	globals []object.Object
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constant:     bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
		globals:      make([]object.Object, GlobalSize),
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
		case code.OPNULL:
			err := vm.push(Null)
			if err != nil {
				return err
			}
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2
			err := vm.push(vm.constant[constIndex])
			if err != nil {
				return err
			}

		case code.OPMINUS:
			err := vm.executeMinusOperator()
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
		case code.OPTRUE:
			err := vm.push(True)
			if err != nil {
				return err
			}
		case code.OPFALSE:
			err := vm.push(False)
			if err != nil {
				return err
			}
		case code.OPEQUAL, code.OPNOTEQUAL, code.OPGREATERTHAN:
			err := vm.executeComparison(op)
			if err != nil {
				return err
			}
		case code.OPBANG:
			err := vm.executeBangOperator()
			if err != nil {
				return err
			}
		case code.OPJUMP:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip = pos - 1

		case code.OPJUMPNOTTRUE:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			condition := vm.pop()
			if !isTruthy(condition) {
				ip = pos - 1
			}

		case code.OPSETGLOBAL:
			globalIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2
			vm.globals[globalIndex] = vm.pop()

		case code.OPGETGLOBAL:
			globalIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			err := vm.push(vm.globals[globalIndex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value
	case *object.Null:
		return false
	default:
		return true
	}
}

func (vm *VM) executeMinusOperator() error {
	operand := vm.pop()
	if operand.Type() != object.INTEGER_OBJ {
		return fmt.Errorf("unsupported type for negation : %s", operand.Type())
	}

	value := operand.(*object.Integer).Value
	return vm.push(&object.Integer{Value: -value})

}
func (vm *VM) executeBangOperator() error {
	operand := vm.pop()

	switch operand {
	case True:
		return vm.push(False)
	case False:
		return vm.push(True)
	case Null:
		return vm.push(True)
	default:
		return vm.push(False)
	}
}
func (vm *VM) executeComparison(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return vm.executeIntegerComparision(op, left, right)
	}

	switch op {
	case code.OPEQUAL:
		return vm.push(nativeBoolToBooleanObject(right == left))
	case code.OPNOTEQUAL:
		return vm.push(nativeBoolToBooleanObject(right != left))
	default:
		return fmt.Errorf("unknown operator : %d ( %s %s)", op, left.Type(), right.Type())
	}
}

func (vm *VM) executeIntegerComparision(
	op code.Opcode,
	left, right object.Object,
) error {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch op {
	case code.OPEQUAL:
		return vm.push(nativeBoolToBooleanObject(rightValue == leftValue))
	case code.OPNOTEQUAL:
		return vm.push(nativeBoolToBooleanObject(rightValue != leftValue))
	case code.OPGREATERTHAN:
		return vm.push(nativeBoolToBooleanObject(leftValue > rightValue))
	default:
		return fmt.Errorf("unknown operator : %d", op)
	}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
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
