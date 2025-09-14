package ast

import "fmt"

// SpreadExpr represents a spread expression.
type SpreadExpr struct {
	Expression ExprNode
	StartPos   int
	EndPos     int
}

// Expr returns the expression of the spread expression.
func (s *SpreadExpr) Expr() string {
	if s.Expression == nil {
		return "..."
	}

	return fmt.Sprintf("...%s", s.Expression.Expr())
}

// StartPosition returns the start position of the spread expression.
func (s *SpreadExpr) StartPosition() int {
	return s.StartPos
}

// EndPosition returns the end position of the spread expression.
func (s *SpreadExpr) EndPosition() int {
	return s.EndPos
}

// Walk walks the spread expression.
func (s *SpreadExpr) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(s)

	if !shouldContinue {
		return
	}

	if s.Expression != nil {
		shouldContinue = fn(s.Expression)

		if !shouldContinue {
			return
		}

		s.Expression.Walk(fn)
	}
}
