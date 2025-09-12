package ast

// AnyLiteral defines a struct for a literal any value.
type AnyLiteral struct {
	Value    any
	StartPos int
	EndPos   int
}

// Expr returns the expression of the any literal.
func (e *AnyLiteral) Expr() string {
	return "any"
}

// StartPosition returns the start position of the any literal.
func (e *AnyLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the any literal.
func (e *AnyLiteral) EndPosition() int {
	return e.EndPos
}

// Walk walks the any literal.
func (e *AnyLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
