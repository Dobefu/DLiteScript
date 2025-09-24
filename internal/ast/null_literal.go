package ast

// NullLiteral defines a struct for a literal null value.
type NullLiteral struct {
	Range Range
}

// Expr returns the expression of the null literal.
func (e *NullLiteral) Expr() string {
	return "null"
}

// GetRange returns the range of the null literal.
func (e *NullLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the null literal.
func (e *NullLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
