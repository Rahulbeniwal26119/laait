package ast

import "go/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Identifier struct {
	Token token.Token
	Value string
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (ls *LetStatement) statementNode()
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (*Identifier) expressionNode()        {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
