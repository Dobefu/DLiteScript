package ast

// BoolLiteral defines a struct for a literal boolean value.
type BoolLiteral struct {
	Value string
	Pos   int
}

// Expr returns the expression of the boolean literal.
func (e *BoolLiteral) Expr() string {
	return e.Value
}

// Position returns the position of the boolean literal.
func (e *BoolLiteral) Position() int {
	return e.Pos
}
