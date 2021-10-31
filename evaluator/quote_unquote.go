package evaluator

import (
	"laait/ast"
	"laait/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
