package ast

import (
	"fmt"
)

// VariableDeclaration represents a variable declaration.
type VariableDeclaration struct {
	Name  string
	Type  string
	Value ExprNode
	Range Range
}

// Expr returns the expression of the variable declaration.
func (v *VariableDeclaration) Expr() string {
	if v.Value == nil {
		return fmt.Sprintf("var %s %s", v.Name, v.Type)
	}

	return fmt.Sprintf("var %s %s = %s", v.Name, v.Type, v.Value.Expr())
}

// GetRange returns the range of the variable declaration.
func (v *VariableDeclaration) GetRange() Range {
	return v.Range
}

// Walk walks the variable declaration and its value.
func (v *VariableDeclaration) Walk(fn func(node ExprNode) bool) {
	shouldContinue := fn(v)

	if !shouldContinue {
		return
	}

	if v.Value != nil {
		shouldContinue = fn(v.Value)

		if !shouldContinue {
			return
		}

		v.Value.Walk(fn)
	}
}
