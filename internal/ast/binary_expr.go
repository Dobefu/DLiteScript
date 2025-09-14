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
	StartPos int
	EndPos   int
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

// StartPosition returns the start position of the binary expression.
func (e *BinaryExpr) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the binary expression.
func (e *BinaryExpr) EndPosition() int {
	return e.EndPos
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
