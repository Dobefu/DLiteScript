package ast

import "fmt"

// StringLiteral defines a struct for a literal string value.
type StringLiteral struct {
	Value string
	Range Range
}

// Expr returns the expression of the string literal.
func (e *StringLiteral) Expr() string {
	return fmt.Sprintf("%q", e.Value)
}

// GetRange returns the range of the string literal.
func (e *StringLiteral) GetRange() Range {
	return e.Range
}

// Walk walks the string literal.
func (e *StringLiteral) Walk(fn func(node ExprNode) bool) {
	fn(e)
}
