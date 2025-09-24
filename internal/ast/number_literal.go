package ast

// NumberLiteral defines a struct for a literal number value.
type NumberLiteral struct {
	Value string
	Range Range
}

// Expr returns the expression of the number literal.
func (e *NumberLiteral) Expr() string {
	return e.Value
}

// GetRange returns the range of the number literal.
func (e *NumberLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the number literal.
func (e *NumberLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
