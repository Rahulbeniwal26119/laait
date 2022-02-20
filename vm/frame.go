// frame is the another name for call stack

package vm

import (
	"laait/code"
	"laait/object"
)

type Frame struct {
	fn          *object.CompiledFunction
	ip          int
	basePointer int // point to bottom of the call stack
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	return &Frame{fn: fn, ip: -1, basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
