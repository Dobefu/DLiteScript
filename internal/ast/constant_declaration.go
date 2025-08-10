package ast

import (
	"fmt"
)

// ConstantDeclaration represents a constant declaration.
type ConstantDeclaration struct {
	Name  string
	Type  string
	Value ExprNode
	Pos   int
}

// Expr returns the expression of the constant declaration.
func (c *ConstantDeclaration) Expr() string {
	if c.Value == nil {
		return fmt.Sprintf("const %s %s", c.Name, c.Type)
	}

	return fmt.Sprintf("const %s %s = %s", c.Name, c.Type, c.Value.Expr())
}

// Position returns the position of the constant declaration.
func (c *ConstantDeclaration) Position() int {
	return c.Pos
}
