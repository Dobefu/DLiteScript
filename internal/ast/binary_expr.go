package ast

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/token"
)

// BinaryExpr defines a struct for a binary expression.
type BinaryExpr struct {
	Left     ExprNode
	Right    ExprNode
	Operator token.Token
	Range    Range
}

// Expr returns the expression of the binary expression.
func (e *BinaryExpr) Expr() string {
	if e.Left == nil || e.Right == nil {
		return ""
	}

	return fmt.Sprintf(
		"(%s %s %s)",
		e.Left.Expr(),
		e.Operator.Atom,
		e.Right.Expr(),
	)
}

// GetRange returns the range of the binary expression.
func (e *BinaryExpr) GetRange() Range {
	return e.Range
}

// Walk walks the binary expression and its left and right nodes.
func (e *BinaryExpr) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(e)

	if !shouldContinue {
		return
	}

	if e.Left != nil {
		shouldContinue = fn(e.Left)

		if !shouldContinue {
			return
		}

		e.Left.Walk(fn)
	}

	if e.Right != nil {
		shouldContinue = fn(e.Right)

		if !shouldContinue {
			return
		}

		e.Right.Walk(fn)
	}
}
