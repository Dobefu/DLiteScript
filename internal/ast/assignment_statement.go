package ast

import (
	"fmt"
)

// AssignmentStatement represents an assignment statement.
type AssignmentStatement struct {
	Left  *Identifier
	Right ExprNode
	Range Range
}

// Expr returns the expression of the assignment statement.
func (a *AssignmentStatement) Expr() string {
	if a.Left == nil || a.Right == nil {
		return ""
	}

	return fmt.Sprintf("%s = %s", a.Left.Expr(), a.Right.Expr())
}

// GetRange returns the range of the assignment statement.
func (a *AssignmentStatement) GetRange() Range {
	return a.Range
}

// Walk walks the assignment statement and its left and right nodes.
func (a *AssignmentStatement) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(a)

	if !shouldContinue {
		return
	}

	if a.Left != nil {
		shouldContinue = fn(a.Left)

		if !shouldContinue {
			return
		}

		a.Left.Walk(fn)
	}

	if a.Right != nil {
		shouldContinue = fn(a.Right)

		if !shouldContinue {
			return
		}

		a.Right.Walk(fn)
	}
}
