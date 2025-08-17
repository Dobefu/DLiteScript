package ast

import (
	"fmt"
)

// ConstantDeclaration represents a constant declaration.
type ConstantDeclaration struct {
	Name     string
	Type     string
	Value    ExprNode
	StartPos int
	EndPos   int
}

// Expr returns the expression of the constant declaration.
func (c *ConstantDeclaration) Expr() string {
	return fmt.Sprintf("const %s %s = %s", c.Name, c.Type, c.Value.Expr())
}

// StartPosition returns the start position of the constant declaration.
func (c *ConstantDeclaration) StartPosition() int {
	return c.StartPos
}

// EndPosition returns the end position of the constant declaration.
func (c *ConstantDeclaration) EndPosition() int {
	return c.EndPos
}

// Walk walks the constant declaration and its value.
func (c *ConstantDeclaration) Walk(fn func(node ExprNode) bool) {
	fn(c)

	c.Value.Walk(fn)
}
