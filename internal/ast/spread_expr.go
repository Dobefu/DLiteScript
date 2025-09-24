package ast

import "fmt"

// SpreadExpr represents a spread expression.
type SpreadExpr struct {
	Expression ExprNode
	Range      Range
}

// Expr returns the expression of the spread expression.
func (s *SpreadExpr) Expr() string {
	if s.Expression == nil {
		return "..."
	}

	return fmt.Sprintf("...%s", s.Expression.Expr())
}

// GetRange returns the range of the spread expression.
func (s *SpreadExpr) GetRange() Range {
	return s.Range
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
