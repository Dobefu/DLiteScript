package ast

// SpreadExpr represents a spread expression.
type SpreadExpr struct {
	Expression ExprNode
	StartPos   int
	EndPos     int
}

// Expr returns the expression of the spread expression.
func (s *SpreadExpr) Expr() string {
	return "..." + s.Expression.Expr()
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
	fn(s)
	s.Expression.Walk(fn)
}
