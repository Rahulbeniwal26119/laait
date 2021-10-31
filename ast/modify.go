package ast

type ModifierFunction func(Node) Node

func Modify(node Node, modifier ModifierFunction) Node {
	switch node := node.(type) {
	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modifier).(Statement)
		}
	case *ExpressionStatement:
		node.Expression, _ = Modify(node.Expression, modifier).(Expression)
	case *InfixExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *PrefixExpression:
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *IndexExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Index, _ = Modify(node.Index, modifier).(Expression)
	case *IfExpression:
		node.Condition, _ = Modify(node.Condition, modifier).(Expression)
		node.Effect, _ = Modify(node.Effect, modifier).(*BlockStatement)
		if node.Optional != nil {
			node.Optional, _ = Modify(node.Optional, modifier).(*BlockStatement)
		}

	case *BlockStatement:
		for i, _ := range node.Statements {
			node.Statements[i], _ = Modify(node.Statements[i], modifier).(Statement)
		}
	}

	return modifier(node)
}
