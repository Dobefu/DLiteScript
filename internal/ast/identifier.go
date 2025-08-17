package ast

// Identifier defines a struct for an identifier.
type Identifier struct {
	Value    string
	StartPos int
	EndPos   int
}

// Expr returns the expression of the identifier.
func (e *Identifier) Expr() string {
	return e.Value
}

// StartPosition returns the start position of the identifier.
func (e *Identifier) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the identifier.
func (e *Identifier) EndPosition() int {
	return e.EndPos
}

// Walk walks the identifier.
func (e *Identifier) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
