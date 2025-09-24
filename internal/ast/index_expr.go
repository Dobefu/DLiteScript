package ast

import (
	"fmt"
)

// IndexExpr defines a struct for an index expression.
type IndexExpr struct {
	Array ExprNode
	Index ExprNode
	Range Range
}

// Expr returns the expression of the index expression.
func (e *IndexExpr) Expr() string {
	if e.Array == nil || e.Index == nil {
		return ""
	}

	return fmt.Sprintf("%s[%s]", e.Array.Expr(), e.Index.Expr())
}

// GetRange returns the range of the index expression.
func (e *IndexExpr) GetRange() Range {
	return e.Range
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
