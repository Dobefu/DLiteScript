package ast

import (
	"fmt"
)

// ContinueStatement represents a continue statement.
type ContinueStatement struct {
	Count int
	Pos   int
}

// Expr returns the expression of the continue statement.
func (c *ContinueStatement) Expr() string {
	if c.Count == 1 {
		return "continue"
	}

	return fmt.Sprintf("continue %d", c.Count)
}

// Position returns the position of the continue statement.
func (c *ContinueStatement) Position() int {
	return c.Pos
}
