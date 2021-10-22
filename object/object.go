package object

import (
	"bytes"
	"fmt"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILT_OBJ        = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
)

type Error struct {
	Message string
}

type ReturnValue struct {
	Value Object
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type String struct {
	Value string
}

type Null struct{}

type Builtin struct {
	Fn BuiltinFunction
}

type BuiltinFunction func(args ...Object) Object

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ","))
	out.WriteString("]")

	return out.String()
}
func (b *Builtin) Type() ObjectType {
	return BUILT_OBJ
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}

func (n *Null) Type() ObjectType { return NULL_OBJ }

func (n *Null) Inspect() string { return "null" }

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (s *String) Type() ObjectType { return STRING_OBJ }

func (s *String) Inspect() string { return s.Value }

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
