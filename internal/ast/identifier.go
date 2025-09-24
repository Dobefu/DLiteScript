package ast

// Identifier defines a struct for an identifier.
type Identifier struct {
	Value string
	Range Range
}

// Expr returns the expression of the identifier.
func (e *Identifier) Expr() string {
	return e.Value
}

// GetRange returns the range of the identifier.
func (e *Identifier) GetRange() Range {
	return e.Range
}

// Walk walks the identifier.
func (e *Identifier) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
