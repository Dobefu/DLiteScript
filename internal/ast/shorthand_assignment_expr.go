package ast

import (
	"fmt"

	"github.com/Dobefu/DLiteScript/internal/token"
)

// ShorthandAssignmentExpr represents a shorthand assignment expression.
type ShorthandAssignmentExpr struct {
	Left     ExprNode
	Right    ExprNode
	Operator token.Token
	Range    Range
}

// Expr returns the expression of the shorthand assignment expression.
func (s *ShorthandAssignmentExpr) Expr() string {
	if s.Left == nil || s.Right == nil {
		return ""
	}

	return fmt.Sprintf(
		"%s %s %s",
		s.Left.Expr(),
		s.Operator.Atom,
		s.Right.Expr(),
	)
}

// GetRange returns the range of the shorthand assignment expression.
func (s *ShorthandAssignmentExpr) GetRange() Range {
	return s.Range
}

// Walk walks the shorthand assignment expreession.
func (s *ShorthandAssignmentExpr) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(s)

	if !shouldContinue {
		return
	}

	if s.Left != nil {
		shouldContinue := fn(s.Left)

		if !shouldContinue {
			return
		}

		s.Left.Walk(fn)
	}

	if s.Right != nil {
		shouldContinue := fn(s.Right)

		if !shouldContinue {
			return
		}

		s.Right.Walk(fn)
	}
}
