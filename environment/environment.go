package environment

import (
	"bytes"
	"laait/ast"
	"laait/object"
	"strings"
)

type Environment struct {
	store map[string]object.Object
	outer *Environment
}

func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val object.Object) object.Object {
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]object.Object)
	return &Environment{store: s, outer: nil}
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() object.ObjectType { return object.FUNCTION_OBJ }

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("function")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type Macro struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (m *Macro) Type() object.ObjectType { return object.MACRO_OBJ }
func (m *Macro) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("macro")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")

	return out.String()

}
