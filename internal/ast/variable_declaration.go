package ast

import (
	"fmt"
)

type VariableDeclaration struct {
	Name  string
	Type  string
	Value ExprNode
	Pos   int
}

func (v *VariableDeclaration) Expr() string {
	if v.Value == nil {
		return fmt.Sprintf("var %s %s", v.Name, v.Type)
	}

	return fmt.Sprintf("var %s %s = %s", v.Name, v.Type, v.Value.Expr())
}

func (v *VariableDeclaration) Position() int {
	return v.Pos
}
