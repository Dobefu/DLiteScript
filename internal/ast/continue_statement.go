package ast

import (
	"fmt"
)

// ContinueStatement represents a continue statement.
type ContinueStatement struct {
	Count    int
	StartPos int
	EndPos   int
}

// Expr returns the expression of the continue statement.
func (c *ContinueStatement) Expr() string {
	if c.Count == 1 {
		return "continue"
	}

	return fmt.Sprintf("continue %d", c.Count)
}

// StartPosition returns the start position of the continue statement.
func (c *ContinueStatement) StartPosition() int {
	return c.StartPos
}

// EndPosition returns the end position of the continue statement.
func (c *ContinueStatement) EndPosition() int {
	return c.EndPos
}

// Walk walks the continue statement.
func (c *ContinueStatement) Walk(fn func(node ExprNode) bool) {
	fn(c)
}
