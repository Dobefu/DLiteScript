package ast

import (
	"fmt"
)

// IndexAssignmentStatement represents an assignment to an array index.
type IndexAssignmentStatement struct {
	Array    ExprNode
	Index    ExprNode
	Right    ExprNode
	StartPos int
	EndPos   int
}

// Expr returns the expression of the index assignment statement.
func (a *IndexAssignmentStatement) Expr() string {
	if a.Array == nil || a.Index == nil || a.Right == nil {
		return ""
	}

	return fmt.Sprintf(
		"%s[%s] = %s",
		a.Array.Expr(),
		a.Index.Expr(),
		a.Right.Expr(),
	)
}

// StartPosition returns the start position of the index assignment statement.
func (a *IndexAssignmentStatement) StartPosition() int {
	return a.StartPos
}

// EndPosition returns the end position of the index assignment statement.
func (a *IndexAssignmentStatement) EndPosition() int {
	return a.EndPos
}

// Walk walks the index assignment statement and its array, index, and right nodes.
func (a *IndexAssignmentStatement) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(a)

	if !shouldContinue {
		return
	}

	if a.Array != nil {
		shouldContinue = fn(a.Array)

		if !shouldContinue {
			return
		}

		a.Array.Walk(fn)
	}

	if a.Index != nil {
		shouldContinue = fn(a.Index)

		if !shouldContinue {
			return
		}

		a.Index.Walk(fn)
	}

	if a.Right != nil {
		a.Right.Walk(fn)
	}
}
