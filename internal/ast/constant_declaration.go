package ast

import (
	"fmt"
)

type ConstantDeclaration struct {
	Name  string
	Type  string
	Value ExprNode
	Pos   int
}

func (c *ConstantDeclaration) Expr() string {
	if c.Value == nil {
		return fmt.Sprintf("const %s %s", c.Name, c.Type)
	}

	return fmt.Sprintf("const %s %s = %s", c.Name, c.Type, c.Value.Expr())
}

func (c *ConstantDeclaration) Position() int {
	return c.Pos
}
