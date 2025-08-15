package ast

import (
	"fmt"
)

// BreakStatement represents a break statement.
type BreakStatement struct {
	Count int
	Pos   int
}

// Expr returns the expression of the break statement.
func (b *BreakStatement) Expr() string {
	if b.Count == 1 {
		return "break"
	}

	return fmt.Sprintf("break %d", b.Count)
}

// Position returns the position of the break statement.
func (b *BreakStatement) Position() int {
	return b.Pos
}
