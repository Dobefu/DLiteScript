package ast

import (
	"fmt"
)

// IfStatement defines a struct for an if statement.
type IfStatement struct {
	Condition ExprNode
	ThenBlock *BlockStatement
	ElseBlock *BlockStatement
	Pos       int
}

// Expr returns the expression of the if statement.
func (e *IfStatement) Expr() string {
	if e.ElseBlock == nil {
		return fmt.Sprintf("if %s { %s }", e.Condition.Expr(), e.ThenBlock.Expr())
	}

	return fmt.Sprintf("if %s { %s } else { %s }", e.Condition.Expr(), e.ThenBlock.Expr(), e.ElseBlock.Expr())
}

// Position returns the position of the if statement.
func (e *IfStatement) Position() int {
	return e.Pos
}
