package ast

import (
	"fmt"
)

// ForStatement represents a for statement.
type ForStatement struct {
	Condition        ExprNode
	Body             *BlockStatement
	Pos              int
	DeclaredVariable string
	RangeVariable    string
	RangeFrom        ExprNode
	RangeTo          ExprNode
	IsRange          bool
}

// Expr returns the expression of the for statement.
func (f *ForStatement) Expr() string {
	if f.Condition == nil {
		return fmt.Sprintf("for { %s }", f.Body.Expr())
	}

	if f.DeclaredVariable != "" {
		return fmt.Sprintf("for var %s %s { %s }", f.DeclaredVariable, f.Condition.Expr(), f.Body.Expr())
	}

	return fmt.Sprintf("for %s { %s }", f.Condition.Expr(), f.Body.Expr())
}

// Position returns the position of the for statement.
func (f *ForStatement) Position() int {
	return f.Pos
}
