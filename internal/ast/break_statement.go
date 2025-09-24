package ast

import (
	"fmt"
)

// BreakStatement represents a break statement.
type BreakStatement struct {
	Count int
	Range Range
}

// Expr returns the expression of the break statement.
func (b *BreakStatement) Expr() string {
	if b.Count == 1 {
		return "break"
	}

	return fmt.Sprintf("break %d", b.Count)
}

// GetRange returns the range of the break statement.
func (b *BreakStatement) GetRange() Range {
	return b.Range
}

// Walk walks the break statement.
func (b *BreakStatement) Walk(fn func(node ExprNode) bool) {
	fn(b)
}
