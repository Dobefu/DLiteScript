package ast

// Identifier defines a struct for an identifier.
type Identifier struct {
	Value string
	Pos   int
}

// Expr returns the expression of the identifier.
func (e *Identifier) Expr() string {
	return e.Value
}

// Position returns the position of the identifier.
func (e *Identifier) Position() int {
	return e.Pos
}
