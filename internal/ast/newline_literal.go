package ast

// NewlineLiteral defines a struct for a literal newline value.
type NewlineLiteral struct {
	StartPos int
	EndPos   int
}

// Expr returns the expression of the newline literal.
func (e *NewlineLiteral) Expr() string {
	return "\n"
}

// StartPosition returns the start position of the newline literal.
func (e *NewlineLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the newline literal.
func (e *NewlineLiteral) EndPosition() int {
	return e.EndPos
}

// Walk walks the newline literal.
func (e *NewlineLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
