package compiler

import (
	"fmt"
	"laait/ast"
	"laait/code"
	"laait/object"
)

type EmittedInstruction struct {
	Opcode   code.Opcode
	Position int
}
type Compiler struct {
	instructions        code.Instructions
	constants           []object.Object
	lastInstruction     EmittedInstruction
	previousInstruction EmittedInstruction
	symbolTable         *SymbolTable
}

type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions:        code.Instructions{},
		constants:           []object.Object{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
		symbolTable:         NewSymbolTable(),
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {

	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		c.emit(code.OPPOP)

	case *ast.IfExpression:
		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}

		jumpNotTruthyPos := c.emit(code.OPJUMPNOTTRUE, 9999)

		err = c.Compile(node.Effect)
		if err != nil {
			return err
		}

		if c.lastInstructionIsPop() {
			c.removeLastPop()
		}

		jumpPos := c.emit(code.OPJUMP, 9999)

		afterEffectPos := len(c.instructions)
		c.changeOperands(jumpNotTruthyPos, afterEffectPos)

		if node.Optional == nil {
			c.emit(code.OPNULL)
		} else {
			err := c.Compile(node.Optional)
			if err != nil {
				return err
			}

			if c.lastInstructionIsPop() {
				c.removeLastPop()
			}
		}

		afterOptionalPos := len(c.instructions)
		c.changeOperands(jumpPos, afterOptionalPos)

	case *ast.BlockStatement:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}
	case *ast.PrefixExpression:
		err := c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "!":
			c.emit(code.OPBANG)
		case "-":
			c.emit(code.OPMINUS)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}
	case *ast.InfixExpression:
		if node.Operator == "<" {
			err := c.Compile(node.Right)
			if err != nil {
				return err
			}

			err = c.Compile(node.Left)
			if err != nil {
				return err
			}
			c.emit(code.OPGREATERTHAN)
			return nil
		}

		err := c.Compile(node.Left)
		if err != nil {
			return err
		}

		err = c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "+":
			c.emit(code.OPADD)
		case "-":
			c.emit(code.OPSUB)
		case "*":
			c.emit(code.OPMUL)
		case "/":
			c.emit(code.OPDIV)
		case ">":
			c.emit(code.OPGREATERTHAN)
		case "==":
			c.emit(code.OPEQUAL)
		case "!=":
			c.emit(code.OPNOTEQUAL)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}

	case *ast.StringLiteral:
		str := &object.String{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(str))

	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))

	case *ast.Boolean:
		if node.Value {
			c.emit(code.OPTRUE)
		} else {
			c.emit(code.OPFALSE)
		}

	case *ast.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			return fmt.Errorf("undefined variable %s", node.Value)
		}

		c.emit(code.OPGETGLOBAL, symbol.Index)

	case *ast.LetStatement:
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}
		symbol := c.symbolTable.Define(node.Name.Value)
		c.emit(code.OPSETGLOBAL, symbol.Index)
	}

	return nil
}

func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)
	c.setLastInstruction(op, pos)
	return pos
}

func (c *Compiler) setLastInstruction(op code.Opcode, pos int) {
	previous := c.lastInstruction
	last := EmittedInstruction{op, pos}
	c.previousInstruction = previous
	c.lastInstruction = last
}

func (c *Compiler) addInstruction(ins []byte) int {
	newInstructionPosition := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return newInstructionPosition
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

func (c *Compiler) lastInstructionIsPop() bool {
	return c.lastInstruction.Opcode == code.OPPOP
}

func (c *Compiler) removeLastPop() {
	c.instructions = c.instructions[:c.lastInstruction.Position]
	c.lastInstruction = c.previousInstruction
}

func (c *Compiler) replaceInstruction(pos int, newInstructions []byte) {
	for i := 0; i < len(newInstructions); i++ {
		c.instructions[pos+i] = newInstructions[i]
	}
}

func (c *Compiler) changeOperands(opPos int, operands ...int) {
	op := code.Opcode(c.instructions[opPos])
	newInstruction := code.Make(op, operands...)
	c.replaceInstruction(opPos, newInstruction)
}

// for repl each time new compiler and vm instance are created
// so for supporting identifier we have to add a state to
// remember symbol during repl

func NewWithState(s *SymbolTable, constants []object.Object) *Compiler {
	compiler := New()
	compiler.symbolTable = s
	compiler.constants = constants
	return compiler
}
