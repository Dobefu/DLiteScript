package ast

import (
	"fmt"
)

// IfStatement defines a struct for an if statement.
type IfStatement struct {
	Condition ExprNode
	ThenBlock *BlockStatement
	ElseBlock *BlockStatement
	StartPos  int
	EndPos    int
}

// Expr returns the expression of the if statement.
func (e *IfStatement) Expr() string {
	if e.ElseBlock == nil {
		return fmt.Sprintf("if %s { %s }", e.Condition.Expr(), e.ThenBlock.Expr())
	}

	return fmt.Sprintf("if %s { %s } else { %s }", e.Condition.Expr(), e.ThenBlock.Expr(), e.ElseBlock.Expr())
}

// StartPosition returns the start position of the if statement.
func (e *IfStatement) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the if statement.
func (e *IfStatement) EndPosition() int {
	return e.EndPos
}
