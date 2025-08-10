package ast

import (
	"fmt"
)

// VariableDeclaration represents a variable declaration.
type VariableDeclaration struct {
	Name  string
	Type  string
	Value ExprNode
	Pos   int
}

// Expr returns the expression of the variable declaration.
func (v *VariableDeclaration) Expr() string {
	if v.Value == nil {
		return fmt.Sprintf("var %s %s", v.Name, v.Type)
	}

	return fmt.Sprintf("var %s %s = %s", v.Name, v.Type, v.Value.Expr())
}

// Position returns the position of the variable declaration.
func (v *VariableDeclaration) Position() int {
	return v.Pos
}
