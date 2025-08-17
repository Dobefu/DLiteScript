package ast

// NumberLiteral defines a struct for a literal number value.
type NumberLiteral struct {
	Value    string
	StartPos int
	EndPos   int
}

// Expr returns the expression of the number literal.
func (e *NumberLiteral) Expr() string {
	return e.Value
}

// StartPosition returns the start position of the number literal.
func (e *NumberLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the number literal.
func (e *NumberLiteral) EndPosition() int {
	return e.EndPos
}
