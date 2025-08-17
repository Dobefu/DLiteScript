package ast

import (
	"fmt"
)

// AssignmentStatement represents an assignment statement.
type AssignmentStatement struct {
	Left     *Identifier
	Right    ExprNode
	StartPos int
	EndPos   int
}

// Expr returns the expression of the assignment statement.
func (a *AssignmentStatement) Expr() string {
	return fmt.Sprintf("%s = %s", a.Left.Expr(), a.Right.Expr())
}

// StartPosition returns the start position of the assignment statement.
func (a *AssignmentStatement) StartPosition() int {
	return a.StartPos
}

// EndPosition returns the end position of the assignment statement.
func (a *AssignmentStatement) EndPosition() int {
	return a.EndPos
}
