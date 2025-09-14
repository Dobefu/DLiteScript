package ast

import (
	"fmt"
)

// IndexExpr defines a struct for an index expression.
type IndexExpr struct {
	Array    ExprNode
	Index    ExprNode
	StartPos int
	EndPos   int
}

// Expr returns the expression of the index expression.
func (e *IndexExpr) Expr() string {
	if e.Array == nil || e.Index == nil {
		return ""
	}

	return fmt.Sprintf("%s[%s]", e.Array.Expr(), e.Index.Expr())
}

// StartPosition returns the start position of the index expression.
func (e *IndexExpr) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the index expression.
func (e *IndexExpr) EndPosition() int {
	return e.EndPos
}

// Walk walks the index expression and its array and index nodes.
func (e *IndexExpr) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(e)

	if !shouldContinue {
		return
	}

	if e.Array != nil {
		shouldContinue = fn(e.Array)

		if !shouldContinue {
			return
		}

		e.Array.Walk(fn)
	}

	if e.Index != nil {
		shouldContinue = fn(e.Index)

		if !shouldContinue {
			return
		}

		e.Index.Walk(fn)
	}
}
