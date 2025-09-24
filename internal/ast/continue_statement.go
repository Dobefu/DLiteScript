package ast

import (
	"fmt"
)

// ContinueStatement represents a continue statement.
type ContinueStatement struct {
	Count int
	Range Range
}

// Expr returns the expression of the continue statement.
func (c *ContinueStatement) Expr() string {
	if c.Count == 1 {
		return "continue"
	}

	return fmt.Sprintf("continue %d", c.Count)
}

// GetRange returns the range of the continue statement.
func (c *ContinueStatement) GetRange() Range {
	return c.Range
}

// Walk walks the continue statement.
func (c *ContinueStatement) Walk(fn func(node ExprNode) bool) {
	fn(c)
}
