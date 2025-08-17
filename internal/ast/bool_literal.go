package ast

// BoolLiteral defines a struct for a literal boolean value.
type BoolLiteral struct {
	Value    string
	StartPos int
	EndPos   int
}

// Expr returns the expression of the boolean literal.
func (e *BoolLiteral) Expr() string {
	return e.Value
}

// StartPosition returns the start position of the boolean literal.
func (e *BoolLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the boolean literal.
func (e *BoolLiteral) EndPosition() int {
	return e.EndPos
}

// Walk walks the boolean literal.
func (e *BoolLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
