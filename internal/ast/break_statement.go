package ast

import (
	"fmt"
)

// BreakStatement represents a break statement.
type BreakStatement struct {
	Count    int
	StartPos int
	EndPos   int
}

// Expr returns the expression of the break statement.
func (b *BreakStatement) Expr() string {
	if b.Count == 1 {
		return "break"
	}

	return fmt.Sprintf("break %d", b.Count)
}

// StartPosition returns the start position of the break statement.
func (b *BreakStatement) StartPosition() int {
	return b.StartPos
}

// EndPosition returns the end position of the break statement.
func (b *BreakStatement) EndPosition() int {
	return b.EndPos
}
