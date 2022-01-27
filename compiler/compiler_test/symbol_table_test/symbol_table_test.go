package compiler_test

import (
	"laait/compiler"
	"testing"
)

func TestDefine(t *testing.T) {
	expected := map[string]compiler.Symbol{
		"a": compiler.Symbol{Name: "a", Scope: compiler.GlobalScope, Index: 0},
		"b": compiler.Symbol{Name: "b", Scope: compiler.GlobalScope, Index: 1},
	}

	global := compiler.NewSymbolTable()

	a := global.Define("a")
	if a != expected["a"] {
		t.Errorf("expected a=%+v, got=%+v", expected["a"], a)
	}

	b := global.Define("b")
	if b != expected["b"] {
		t.Errorf("expected b=%+v, got=%+v", expected["b"], b)
	}
}

func TestResolveGlobal(t *testing.T) {
	global := compiler.NewSymbolTable()
	global.Define("a")
	global.Define("b")

	expected := []compiler.Symbol{
		compiler.Symbol{Name: "a", Scope: compiler.GlobalScope, Index: 0},
		compiler.Symbol{Name: "b", Scope: compiler.GlobalScope, Index: 1},
	}

	for _, sym := range expected {
		result, ok := global.Resolve(sym.Name)
		if !ok {
			t.Errorf("expected %s to resolve to %+v, got=%+v", sym.Name, sym, result)
		}
	}
}
