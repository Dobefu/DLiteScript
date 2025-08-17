package ast

import "fmt"

// StringLiteral defines a struct for a literal string value.
type StringLiteral struct {
	Value    string
	StartPos int
	EndPos   int
}

// Expr returns the expression of the string literal.
func (e *StringLiteral) Expr() string {
	return fmt.Sprintf("\"%s\"", e.Value)
}

// StartPosition returns the start position of the string literal.
func (e *StringLiteral) StartPosition() int {
	return e.StartPos
}

// EndPosition returns the end position of the string literal.
func (e *StringLiteral) EndPosition() int {
	return e.EndPos
}

// Walk walks the string literal.
func (e *StringLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
