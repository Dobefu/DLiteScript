package ast

// NullLiteral defines a struct for a literal null value.
type NullLiteral struct {
	Pos int
}

// Expr returns the expression of the null literal.
func (e *NullLiteral) Expr() string {
	return "null"
}

// Position returns the position of the null literal.
func (e *NullLiteral) Position() int {
	return e.Pos
}
