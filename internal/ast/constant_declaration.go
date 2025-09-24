package ast

import (
	"fmt"
)

// ConstantDeclaration represents a constant declaration.
type ConstantDeclaration struct {
	Name  string
	Type  string
	Value ExprNode
	Range Range
}

// Expr returns the expression of the constant declaration.
func (c *ConstantDeclaration) Expr() string {
	if c.Value == nil {
		return ""
	}

	return fmt.Sprintf("const %s %s = %s", c.Name, c.Type, c.Value.Expr())
}

// GetRange returns the range of the constant declaration.
func (c *ConstantDeclaration) GetRange() Range {
	return c.Range
}

// Walk walks the constant declaration and its value.
func (c *ConstantDeclaration) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(c)

	if !shouldContinue {
		return
	}

	if c.Value != nil {
		shouldContinue = fn(c.Value)

		if !shouldContinue {
			return
		}

		c.Value.Walk(fn)
	}
}
