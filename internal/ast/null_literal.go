package ast

// NullLiteral defines a struct for a literal null value.
type NullLiteral struct {
	StartPos int
	EndPos   int
}

// Expr returns the expression of the null literal.
func (e *NullLiteral) Expr() string {
	return "null"
}

// StartPosition returns the start position of the null literal.
func (e *NullLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the null literal.
func (e *NullLiteral) EndPosition() int {
	return e.EndPos
}

// Walk walks the null literal.
func (e *NullLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
