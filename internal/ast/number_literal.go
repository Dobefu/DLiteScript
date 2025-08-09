package ast

// NumberLiteral defines a struct for a literal number value.
type NumberLiteral struct {
	Value string
	Pos   int
}

// Expr returns the expression of the number literal.
func (e *NumberLiteral) Expr() string {
	return e.Value
}

// Position returns the position of the number literal.
func (e *NumberLiteral) Position() int {
	return e.Pos
}
