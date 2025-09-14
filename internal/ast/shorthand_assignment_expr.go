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
	StartPos int
	EndPos   int
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

// StartPosition returns the start position of the shorthand assignment expression.
func (s *ShorthandAssignmentExpr) StartPosition() int {
	return s.StartPos
}

// EndPosition returns the end position of the shorthand assignment expression.
func (s *ShorthandAssignmentExpr) EndPosition() int {
	return s.EndPos
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
