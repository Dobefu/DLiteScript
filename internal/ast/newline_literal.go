package ast

// NewlineLiteral defines a struct for a literal newline value.
type NewlineLiteral struct {
	Range Range
}

// Expr returns the expression of the newline literal.
func (e *NewlineLiteral) Expr() string {
	return "\n"
}

// GetRange returns the range of the newline literal.
func (e *NewlineLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the newline literal.
func (e *NewlineLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
