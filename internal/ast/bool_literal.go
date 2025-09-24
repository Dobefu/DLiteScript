package ast

// BoolLiteral defines a struct for a literal boolean value.
type BoolLiteral struct {
	Value string
	Range Range
}

// Expr returns the expression of the boolean literal.
func (e *BoolLiteral) Expr() string {
	return e.Value
}

// GetRange returns the range of the boolean literal.
func (e *BoolLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the boolean literal.
func (e *BoolLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
