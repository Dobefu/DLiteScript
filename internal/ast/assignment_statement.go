package ast

import (
	"fmt"
)

// AssignmentStatement represents an assignment statement.
type AssignmentStatement struct {
	Left  *Identifier
	Right ExprNode
	Pos   int
}

// Expr returns the expression of the assignment statement.
func (a *AssignmentStatement) Expr() string {
	return fmt.Sprintf("%s = %s", a.Left.Expr(), a.Right.Expr())
}

// Position returns the position of the assignment statement.
func (a *AssignmentStatement) Position() int {
	return a.Pos
}
