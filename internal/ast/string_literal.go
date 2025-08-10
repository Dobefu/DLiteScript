package ast

// StringLiteral defines a struct for a literal string value.
type StringLiteral struct {
	Value string
	Pos   int
}

// Expr returns the expression of the string literal.
func (e *StringLiteral) Expr() string {
	return e.Value
}

// Position returns the position of the string literal.
func (e *StringLiteral) Position() int {
	return e.Pos
}
